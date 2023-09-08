package models

import (
	"time"
)

type User struct {
	ID        int64  `gorm:"primaryKey;column:id" json:"id"`
	FirstName string `gorm:"column:firstName;not null" json:"first_name"`
	LastName  string `gorm:"column:lastName" json:"last_name"`
	Email     string `gorm:"column:email;not null" json:"email"`
	// Password    string    `gorm:"->;column:password;" json:"password,omitempty"`
	Password string `gorm:"->;column:password;" json:"password,omitempty"`

	Active     bool      `gorm:"column:active" json:"active"`
	IsVerified bool      `gorm:"column:isVerified" json:"is_verified"`
	Role       string    `gorm:"column:role;not null" json:"role,omitempty"`
	CreatedOn  time.Time `gorm:"column:createdOn" json:"created_on"`
	Phone      string    `gorm:"column:phone;not null" json:"phone,omitempty"`
	//AccessToken field without a corresponding gorm tag
	AccessToken string `gorm:"-" json:"access_token,omitempty"`
}
