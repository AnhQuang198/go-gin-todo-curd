package business

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/entity"
	"social-todo-list/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *entity.TodoItem) error
}

type createItemBusiness struct {
	store CreateItemStorage
}

func NewCreateItemBusiness(store CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{store: store}
}

func (business *createItemBusiness) CreateNewItem(ctx context.Context, req *model.TodoItemRequest) error {
	title := strings.TrimSpace(req.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	item := req.ToItemEntity()
	if err := business.store.CreateItem(ctx, item); err != nil {
		return common.ErrCannotCreateEntity("Item", err)
	}
	return nil
}
