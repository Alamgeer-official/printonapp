package models

import (
	"time"
)

type UserRole string

const (
	UserRoleAdmin    UserRole = "ADMIN"
	UserRoleCustomer UserRole = "USER"
)

type User struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	FirstName  string    `gorm:"column:firstName;not null" json:"first_name"`
	LastName   string    `gorm:"column:lastName" json:"last_name"`
	Email      string    `gorm:"column:email;not null" json:"email"`
	Password   string    `gorm:"column:password;not null" json:"password,omitempty"`
	Active     bool      `gorm:"column:active" json:"active"`
	IsVerified bool      `gorm:"column:isVerified" json:"is_verified"`
	Role       UserRole  `gorm:"column:role;not null" json:"role"`
	CreatedOn  time.Time `gorm:"column:createdOn" json:"created_on"`
	Phone      string    `gorm:"column:phone;not null" json:"phone"`
}
