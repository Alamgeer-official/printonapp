package services

import (
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
)

type CollegeSrv interface {
	GetColleges() (*[]models.College, error)
}

type collegeSrv struct {
	collegeRepo repository.CollegeRepo
}

func NewCollegeRepo(cRepo repository.CollegeRepo) CollegeSrv {
	return &collegeSrv{collegeRepo: cRepo}
}

func (cs *collegeSrv) GetColleges() (*[]models.College, error) {
	data, err := cs.collegeRepo.GetColleges()
	if err != nil {
		return nil, err
	}
	return data, err

}
