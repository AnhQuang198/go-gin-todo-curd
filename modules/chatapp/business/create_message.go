package business

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/chatapp/entity"
)

type CreateMessageStorage interface {
	CreateMessage(ctx context.Context, data *entity.Message) error
}

type createMessageBusiness struct {
	store CreateMessageStorage
}

func NewCreateMessageBusiness(store CreateMessageStorage) *createMessageBusiness {
	return &createMessageBusiness{store: store}
}

func (business *createMessageBusiness) CreateNewMessage(ctx context.Context, roomId, senderId, msg string) error {
	msgSave := entity.MappingData(roomId, senderId, msg)
	if err := business.store.CreateMessage(ctx, msgSave); err != nil {
		return common.ErrCannotCreateEntity("Message", err)
	}
	return nil
}
