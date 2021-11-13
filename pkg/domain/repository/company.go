package repository

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type CompanyRepository interface {
	DBGetAll() ([]*entity.Company, error)
	DBCreate(*gin.Context, entity.Company) error
}

type ICompanies interface {
	Get(id string) (entity.Company, error)
	GetAll() ([]entity.Company, error)
	Create() error
}

type Companies struct {
	ID   int
	Name string
}
