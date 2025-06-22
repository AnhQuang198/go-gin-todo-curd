package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (sql *sqlStore) GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := sql.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
