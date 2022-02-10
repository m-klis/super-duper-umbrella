package service

import (
	"gochicoba/models"
	"gochicoba/repository"
)

type ItemService interface {
	GetAllItems() ([]*models.Item, error)
	GetItem(id int) (*models.Item, error)
	// AddItem(item *models.Item) error
	// GetItemById(itemId int) (models.Item, error)
	// DeleteItem(itemId int) error
	// UpdateItem(itemId int, itemData models.Item) (models.Item, error)
}

type itemService struct {
	itemRepo repository.ItemRepository
}

func NewItemService(itemRepo repository.ItemRepository) ItemService {
	return &itemService{
		itemRepo: itemRepo,
	}
}

func (is *itemService) GetAllItems() ([]*models.Item, error) {
	list, err := is.itemRepo.GetAllItems()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (is *itemService) GetItem(id int) (*models.Item, error) {
	item, err := is.itemRepo.GetItem(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}
