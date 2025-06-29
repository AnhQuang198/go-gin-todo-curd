package storage

import (
	"context"
	"social-todo-list/modules/chatapp/entity"
)

func (sql *sqlStore) CreateMessage(ctx context.Context, data *entity.Message) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
