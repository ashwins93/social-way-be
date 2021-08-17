package model

import "time"

var Models []interface{}

func Register(model interface{}) {
	Models = append(Models, model)
}

type Auditable struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
