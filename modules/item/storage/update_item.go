package storage

import (
	"context"
	"social-todo-list/modules/item/entity"
)

func (sql *sqlStore) UpdateItem(ctx context.Context, condition map[string]interface{}, data *entity.TodoItem) error {
	if err := sql.db.Where(condition).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
