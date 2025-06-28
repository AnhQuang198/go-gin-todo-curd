package business

import (
	"context"
	"social-todo-list/modules/item/entity"
	"social-todo-list/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*entity.TodoItem, error)
}

type getItemBusiness struct {
	store GetItemStorage
}

func NewGetItemBusiness(store GetItemStorage) *getItemBusiness {
	return &getItemBusiness{store: store}
}

func (business *getItemBusiness) GetItemById(ctx context.Context, id int) (*model.TodoItemResponse, error) {
	data, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	itemResponse := model.FromEntity(data)
	return itemResponse, nil
}
