package business

import (
	"context"
	"social-todo-list/modules/item/entity"
	"social-todo-list/modules/item/enum"
	"social-todo-list/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*entity.TodoItem, error)
	UpdateItem(ctx context.Context, condition map[string]interface{}, data *entity.TodoItem) error
}

type updateItemBusiness struct {
	store UpdateItemStorage
}

func UpdateItemBusiness(store UpdateItemStorage) *updateItemBusiness {
	return &updateItemBusiness{store: store}
}

func (business *updateItemBusiness) UpdateItemById(ctx context.Context, id int, req *model.TodoItemRequest) error {
	data, err := business.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status != nil && *data.Status == enum.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	itemSave := req.ToItemEntity()
	if err := business.store.UpdateItem(ctx, map[string]interface{}{"id": id}, itemSave); err != nil {
		return err
	}
	return nil
}
