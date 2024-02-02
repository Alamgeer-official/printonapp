package repository

import (
	"errors"

	"githuh.com/printonapp/models"
	"gorm.io/gorm"
)

type ThesisRepo interface {
	CreateThesis(thesis *models.Theses) error
	ReadAllTheses() (*[]models.Theses, error)
	GetThesisByID(id uint64) (*models.Theses, error)
}

type thesisRepo struct{}

func NewThesisRepo() ThesisRepo {
	return &thesisRepo{}
}

func (tr *thesisRepo) CreateThesis(thesis *models.Theses) error {
	if res := gormDB.Create(thesis); res.Error != nil {
		return errors.New("error in creating thesis")
	}
	return nil
}

func (tr *thesisRepo) ReadAllTheses() (*[]models.Theses, error) {
	var theses []models.Theses
	if res := gormDB.Find(&theses); res.Error != nil {
		return nil, errors.New("error in getting thesis records")
	}
	return &theses, nil
}

func (tr *thesisRepo) GetThesisByID(id uint64) (*models.Theses, error) {
	var thesis models.Theses
	if res := gormDB.Where("id = ?", id).First(&thesis); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("thesis not found")
		}
		return nil, errors.New("error in getting thesis by ID")
	}
	return &thesis, nil
}
