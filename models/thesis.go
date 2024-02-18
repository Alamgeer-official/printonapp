package models

import (
	"time"
)

type Theses struct {
	ID           uint64      `gorm:"column:id;primaryKey" json:"id"`
	Active       bool        `gorm:"column:active;not null;default:false" json:"active"`
	CreatedOn    time.Time   `gorm:"column:createdon;not null;default:CURRENT_TIMESTAMP" json:"created_on"`
	UpdatedOn    time.Time   `gorm:"column:updatedon;default:CURRENT_TIMESTAMP" json:"updated_on"`
	CreatedBy    uint64      `gorm:"column:createdby" json:"created_by"`
	User         User        `gorm:"foreignKey:CreatedBy;references:ID" json:"user"`
	UpdatedBy    uint64      `gorm:"column:updatedby" json:"updated_by"`
	Color        string      `gorm:"column:color;not null" json:"color"`
	PaperType    string      `gorm:"column:papertype;not null" json:"paper_type"`
	Description  string      `gorm:"column:description;not null" json:"description"`
	Quantity     uint64      `gorm:"column:quantity;not null" json:"quantity"`
	EstimateCost float64     `gorm:"column:estimatecost;not null" json:"estimate_cost"`
	PDF          string      `gorm:"column:pdf;not null" json:"pdf"`
	Status       OrderStatus `gorm:"column:status;not null" json:"status"`
}

type OrderStatus string

const (
	Booked     OrderStatus = "BOOKED"
	InProgress OrderStatus = "IN PROGRESS"
	Completed  OrderStatus = "COMPLETED"
)
