package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tags struct {
	Id   string `json:"id" gorm:"primary_key;comment:'uuid v4 from service'"`
	Name string `json:"name" gorm:"type:varchar(255)"`
}

// hook function, executed automatically by gorm
func (t *Tags) BeforeCreate(tx *gorm.DB) (err error) {
	t.Id = uuid.NewString()
	return
}
