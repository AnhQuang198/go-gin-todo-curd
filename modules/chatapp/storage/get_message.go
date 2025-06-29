package storage

import (
	"context"
	"social-todo-list/modules/chatapp/entity"
)

func (sql *sqlStore) GetMessage(ctx context.Context) ([]*entity.Message, error) {
	var result []*entity.Message

	if err := sql.db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
