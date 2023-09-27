package entity

import "time"

type Role struct {
	Code        string    `gorm:"primaryKey;column:CODE;size:50" json:"code"`
	Name        string    `gorm:"unique;column:NAME;size:100" json:"name"`
	CreatedDate time.Time `gorm:"column:CREATED_DATE" json:"createdDate"`
	UpdatedDate time.Time `gorm:"column:UPDATED_DATE" json:"updatedDate"`
}

func (Role) TableName() string {
	return "ROLE"
}
