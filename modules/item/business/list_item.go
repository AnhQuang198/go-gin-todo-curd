package business

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Pagging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBusiness struct {
	store ListItemStorage
}

func NewListItemBusiness(store ListItemStorage) *listItemBusiness {
	return &listItemBusiness{store: store}
}

func (business *listItemBusiness) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Pagging,
) ([]model.TodoItem, error) {
	data, err := business.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
