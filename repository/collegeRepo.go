package repository

import (
	"errors"

	"githuh.com/printonapp/models"
)

type CollegeRepo interface {
	GetColleges() (*[]models.College, error)
}

type collegeRepo struct{}

func NewCollegeRepo() CollegeRepo {
	return &collegeRepo{}

}

func (cr *collegeRepo) GetColleges() (*[]models.College, error) {
	var colleges []models.College
	if res := gormDB.Debug().Where("active", true).Order("collegeName ASC",).Find(&colleges); res.Error != nil {
		return nil, errors.New("error in getting college records")
	}
	return &colleges, nil

}
