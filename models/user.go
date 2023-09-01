package models

import "time"

type UserRole string

const (
	UserRoleAdmin    UserRole = "ADMIN"
	UserRoleCustomer UserRole = "USER"
)

type User struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	FirstName  string    `gorm:"column:firstName" json:"first_name"`
	LastName   string    `gorm:"column:lastName" json:"last_name"`
	Email      string    `gorm:"column:email" json:"email"`
	Password   string    `gorm:"column:password" json:"-"`
	Active     bool      `gorm:"column:active" json:"active"`
	IsVerified bool      `gorm:"column:isVerifyied" json:"is_verified"`
	Role       UserRole  `gorm:"column:role" json:"role"`
	CreatedOn  time.Time `gorm:"column:createdOn" json:"created_on"`
}
