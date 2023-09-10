package models

import "time"

type Product struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Active    bool      `gorm:"column:active" json:"active"`
	UserID    int64     `gorm:"column:userid" json:"user_id"` // You can specify the foreign key relationship here
	CreatedOn time.Time `gorm:"column:createdon" json:"created_on"`
}
