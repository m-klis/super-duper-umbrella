package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"gochicoba/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAllItems(models.ItemFilter) ([]*models.Item, error)
	// GetAllItemsByDate(time.Time, time.Time) ([]*models.Item, error)
	GetItem(int) (*models.Item, error)
	AddItem(*models.Item) (*models.Item, error)
	// GetItemById(itemId int) (models.Item, error)
	DeleteItem(int) error
	UpdateItem(int, *models.Item) (*models.Item, error)
}

type itemRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewItemRepository(db *gorm.DB, redis *redis.Client) ItemRepository {
	return &itemRepository{
		db:    db,
		redis: redis,
	}
}

func (ir *itemRepository) GetAllItems(startEnd models.ItemFilter) ([]*models.Item, error) {
	var list []*models.Item
	query := ir.db
	var check string = "data"

	if startEnd.StartDate != nil && startEnd.EndDate != nil {
		query = query.Where("created_at BETWEEN ? AND ?", startEnd.StartDate, startEnd.EndDate)
		check = check + startEnd.StartDate.String() + startEnd.EndDate.String()
	}

	if startEnd.Name != "" {
		query = query.Where("name LIKE ?", "%"+startEnd.Name+"%")
		check = check + startEnd.Name
	}

	if startEnd.Page == 0 {
		startEnd.Page = 1
	}
	switch {
	case startEnd.View > 25:
		startEnd.View = 25
	case startEnd.View <= 0:
		startEnd.View = 5
	}

	offset := (startEnd.Page - 1) * startEnd.View

	check = check + strconv.Itoa(offset)

	str, err := ir.redis.Get(check).Result()
	// fmt.Println(err)

	if err != nil {
		err = query.Order("ID ASC").Offset(offset).Limit(startEnd.View).Find(&list).Error
		if err != nil {
			return nil, err
		}

		var byteList []byte
		byteList, err = json.Marshal(list)
		if err != nil {
			return nil, err
		}

		ir.redis.Set(check, string(byteList), 10*time.Minute)
		fmt.Println("set redis")
	} else {
		resByte := []byte(str)
		// fmt.Println("ResByte :", resByte)
		// fmt.Println("String Redis :", str)
		var data []models.Item
		err = json.Unmarshal(resByte, &data)
		if err != nil {
			return nil, err
		}
		// fmt.Println("Data :", data)
		for i := range data {
			list = append(list, &data[i])
		}
		fmt.Println("get redis")
	}

	return list, nil
}

// func (ir *itemRepository) GetAllItemsByDate(startDate time.Time, endDate time.Time) ([]*models.Item, error) {
// 	var list []*models.Item
// 	query := ir.db
// 	//fmt.Printf("%v\n%v\n", startDate, endDate)
// 	err := query.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&list).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return list, nil
// }

func (ir *itemRepository) GetItem(id int) (item *models.Item, err error) {
	query := ir.db
	err = query.Where("id = ?", id).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (ir *itemRepository) AddItem(itemData *models.Item) (item *models.Item, err error) {
	query := ir.db
	err = query.Create(&itemData).Error
	if err != nil {
		return nil, err
	}
	// double cek
	err = query.Where("id = ?", itemData.ID).First(&item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (ir *itemRepository) UpdateItem(itemId int, itemData *models.Item) (item *models.Item, err error) {
	query := ir.db

	err = query.Model(&itemData).Where("id", itemId).Updates(&itemData).Error
	if err != nil {
		return nil, err
	}
	// double cek
	err = query.Where("id = ?", itemId).First(&itemData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return itemData, nil
}

func (ir *itemRepository) DeleteItem(itemId int) error {
	query := ir.db
	var item *models.Item

	// double cek
	err := query.Where("id = ?", itemId).First(&item).Error
	if err != nil {
		return err
	}

	err = query.Delete(&item).Error
	if err != nil {
		return err
	}

	return nil
}

// func (ir *itemRepository) AddItem(item *models.Item) error {
// 	var id int
// 	var createdAt string
// 	query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at`
// 	err := ir.db.Conn.QueryRow(query, item.Name, item.Description).Scan(&id, &createdAt)
// 	if err != nil {
// 		return err
// 	}
// 	item.ID = id
// 	item.CreatedAt = createdAt
// 	return nil
// }

// func (ir *itemRepository) GetItemById(itemId int) (models.Item, error) {
// 	item := models.Item{}
// 	query := `SELECT * FROM items WHERE id = $1;`
// 	row := ir.db.Conn.QueryRow(query, itemId)
// 	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt); err {
// 	case sql.ErrNoRows:
// 		return item, ErrNoMatch
// 	default:
// 		return item, err
// 	}
// }

// func (ir *itemRepository) DeleteItem(itemId int) error {
// 	query := `DELETE FROM items WHERE id = $1;`
// 	_, err := ir.db.Conn.Exec(query, itemId)
// 	switch err {
// 	case sql.ErrNoRows:
// 		return ErrNoMatch
// 	default:
// 		return err
// 	}
// }

// func (ir *itemRepository) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
// 	item := models.Item{}
// 	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
// 	err := ir.db.Conn.QueryRow(query, itemData.Name, itemData.Description, itemId).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return item, ErrNoMatch
// 		}
// 		return item, err
// 	}
// 	return item, nil
// }
