package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/entity"
	"social-todo-list/modules/item/model"
)

func (sql *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Pagging,
	moreKeys ...string,
) ([]*entity.TodoItem, error) {
	var result []*entity.TodoItem

	db := sql.db.WithContext(ctx).Model(&entity.TodoItem{})

	db = db.Where("status <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
