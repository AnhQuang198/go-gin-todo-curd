package common

import "time"

type SQLModel struct {
	Id       int        `json:"id"`
	CreateAt *time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at,omitempty"`
}
