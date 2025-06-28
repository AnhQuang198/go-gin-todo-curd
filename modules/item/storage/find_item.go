package storage

import (
	"context"
	"social-todo-list/modules/item/entity"
)

func (sql *sqlStore) GetItem(ctx context.Context, condition map[string]interface{}) (*entity.TodoItem, error) {
	var data entity.TodoItem

	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
