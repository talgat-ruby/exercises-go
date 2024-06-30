package todo

import "time"

type Entity struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreateAt    *time.Time `json:"createdAt"`
	UpdateAt    *time.Time `json:"updateAt"`
}
