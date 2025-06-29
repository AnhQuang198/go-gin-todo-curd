package business

import (
	"context"
	"social-todo-list/modules/chatapp/entity"
	"social-todo-list/modules/chatapp/model"
)

type GetMessageStorage interface {
	GetMessage(ctx context.Context) ([]*entity.Message, error)
}

type getMessageBusiness struct {
	store GetMessageStorage
}

func NewGetMessageBusiness(store GetMessageStorage) *getMessageBusiness {
	return &getMessageBusiness{store: store}
}

func (business *getMessageBusiness) GetMessage(ctx context.Context) ([]model.MessageDTO, error) {
	data, err := business.store.GetMessage(ctx)
	if err != nil {
		return nil, err
	}

	return model.FromEntityList(data), nil
}
