package repository

import (
	"errors"
	"gochicoba/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAllItems() ([]*models.Item, error)
	GetItem(id int) (item *models.Item, err error)
	AddItem(item *models.Item) (id int, err error)
	// GetItemById(itemId int) (models.Item, error)
	DeleteItem(itemId int) error
	UpdateItem(itemId int, itemData *models.Item) (item *models.Item, err error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{
		db: db,
	}
}

func (ir *itemRepository) GetAllItems() (itemList []*models.Item, err error) {
	var list []*models.Item
	query := ir.db
	err = query.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

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

func (ir *itemRepository) AddItem(item *models.Item) (id int, err error) {
	query := ir.db
	err = query.Create(&item).Error
	//fmt.Println(item)
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (ir *itemRepository) UpdateItem(itemId int, itemData *models.Item) (item *models.Item, err error) {
	query := ir.db
	//field := &itemData
	err = query.Model(&itemData).Where("id", itemId).Updates(&itemData).Error
	if err != nil {
		return nil, err
	}
	err = query.Where("id = ?", itemId).First(&itemData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return itemData, nil
}

func (ir *itemRepository) DeleteItem(itemId int) error {
	query := ir.db
	var item *models.Item
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
