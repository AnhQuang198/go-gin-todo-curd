package storage

import (
	"context"
	"social-todo-list/modules/item/entity"
)

func (sql *sqlStore) CreateItem(ctx context.Context, data *entity.TodoItem) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
