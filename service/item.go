package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type ItemService interface {
	GetAllItems(models.ItemFilter) ([]*models.Item, error)
	// GetAllItemsByDate(time.Time, time.Time) ([]*models.Item, error)
	GetItem(id int) (*models.Item, error)
	AddItem(item *models.Item) (*models.Item, error)
	// GetItemById(itemId int) (models.Item, error)
	DeleteItem(itemId int) error
	UpdateItem(itemId int, itemData *models.Item) (*models.Item, error)
}

type itemService struct {
	itemRepo repository.ItemRepository
}

func NewItemService(itemRepo repository.ItemRepository) ItemService {
	return &itemService{
		itemRepo: itemRepo,
	}
}

func (is *itemService) GetAllItems(startEnd models.ItemFilter) ([]*models.Item, error) {
	list, err := is.itemRepo.GetAllItems(startEnd)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// func (is *itemService) GetAllItemsByDate(startDate time.Time, endDate time.Time) ([]*models.Item, error) {
// 	list, err := is.itemRepo.GetAllItemsByDate(startDate, endDate)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return list, nil
// }

func (is *itemService) GetItem(id int) (*models.Item, error) {
	item, err := is.itemRepo.GetItem(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (is *itemService) AddItem(item *models.Item) (*models.Item, error) {

	itemData, err := is.itemRepo.AddItem(item)

	if err != nil {
		return nil, err
	}

	return itemData, nil
}

func (is *itemService) UpdateItem(itemId int, itemData *models.Item) (*models.Item, error) {

	item, err := is.itemRepo.UpdateItem(itemId, itemData)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (is *itemService) DeleteItem(itemId int) error {
	err := is.itemRepo.DeleteItem(itemId)

	if err != nil {
		return err
	}

	return nil
}
