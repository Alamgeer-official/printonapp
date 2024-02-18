package repository

import (
	"errors"

	"githuh.com/printonapp/models"
	"gorm.io/gorm"
)

type ThesisRepo interface {
	CreateThesis(thesis *models.Theses) error
	ReadAllTheses() (*[]models.Theses, error)
	ReadAllThesesByUserID(uID int64, page, pageSize int) (*[]models.Theses, int64, error)
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
	if res := gormDB.Preload("User").Find(&theses); res.Error != nil {
		return nil, errors.New("error in getting thesis records")
	}
	return &theses, nil
}

func (tr *thesisRepo) ReadAllThesesByUserID(uID int64, page, pageSize int) (*[]models.Theses, int64, error) {
	var theses []models.Theses
	var totalCount int64

	// Count total number of theses for the user
	if err := gormDB.Model(&models.Theses{}).Where("active = ?", true).Where("createdby = ?", uID).Count(&totalCount).Error; err != nil {
		return nil, 0, errors.New("error counting theses")
	}

	// Retrieve paginated theses for the user
	if res := gormDB.Where("active = ?", true).Where("createdby = ?", uID).Offset((page - 1) * pageSize).Limit(pageSize).Find(&theses); res.Error != nil {
		return nil, 0, errors.New("error fetching thesis records")
	}

	return &theses, totalCount, nil
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
