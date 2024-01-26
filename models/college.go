package models

import (
	"time"
)

type College struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CollegeName string    `gorm:"column:collegeName" json:"collegeName"`
	Active      bool      `gorm:"column:active" json:"active"`
	CreatedOn   time.Time `gorm:"column:createdOn" json:"createdOn"`
}
