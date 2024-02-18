package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/utils"
)

type ThesisSrv interface {
	CreateThesis(ctx *gin.Context, thesis *models.Theses) error
	ReadAllTheses(ctx *gin.Context) (*[]models.Theses, error)
	ReadAllThesesByRole(ctx *gin.Context, collegeID,page, pageSize int) (*models.Pagination, error)
	GetThesisByID(ctx *gin.Context, id uint64) (*models.Theses, error)
}

type thesisSrv struct {
	thesisRepo repository.ThesisRepo
}

func NewThesisSrv(tRepo repository.ThesisRepo) ThesisSrv {
	return &thesisSrv{thesisRepo: tRepo}
}

func (ts *thesisSrv) CreateThesis(ctx *gin.Context, thesis *models.Theses) error {
	user := utils.GetUserDataFromContext(ctx)
	if !user.IsUser() {
		return errors.New("only Authenticated User allow")
	}
	thesis.Active = true
	thesis.CreatedBy = uint64(user.ID)
	thesis.Status = models.OrderStatus(models.Booked)

	if err := ts.thesisRepo.CreateThesis(thesis); err != nil {
		return err
	}
	return nil
}

func (ts *thesisSrv) ReadAllTheses(ctx *gin.Context) (*[]models.Theses, error) {
	user := utils.GetUserDataFromContext(ctx)
	if !user.IsAdmin() {
		return nil, errors.New("only Admin allow")
	}
	data, err := ts.thesisRepo.ReadAllTheses()
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (ts *thesisSrv) ReadAllThesesByRole(ctx *gin.Context,collegeID, page, pageSize int) (*models.Pagination, error) {
	user := utils.GetUserDataFromContext(ctx)
	if user.IsUser() {
		//for normal user
		data, totalCount, err := ts.thesisRepo.ReadAllThesesByUserID(user.ID, page, pageSize)
		if err != nil {
			return nil, err
		}
		pagination := utils.CalculatePagination(totalCount, int64(pageSize), int64(page), data)
		return pagination, nil
	} else if user.IsAdmin() {
		//for admin
		data, totalCount, err := ts.thesisRepo.ReadAllThesesByCollegeID(user.ID,collegeID, page, pageSize)
		if err != nil {
			return nil, err
		}
		pagination := utils.CalculatePagination(totalCount, int64(pageSize), int64(page), data)
		return pagination, nil

	} else {
		return nil, errors.New("Uauthenticated user")
	}
}

func (ts *thesisSrv) GetThesisByID(ctx *gin.Context, id uint64) (*models.Theses, error) {
	user := utils.GetUserDataFromContext(ctx)
	if !user.IsAdmin() && !user.IsUser() {
		return nil, errors.New("only authenticated users and admins are allowed")
	}
	thesis, err := ts.thesisRepo.GetThesisByID(id)
	if err != nil {
		return nil, err
	}
	return thesis, nil
}
