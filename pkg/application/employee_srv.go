package application

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/service"
)

type EmployeeSrv struct {
	repo repository.IEmployee
}

func NewEmployeeSrv(repo repository.IEmployee) service.IEmployee {
	return &EmployeeSrv{
		repo,
	}
}

func (e *EmployeeSrv) Add(employee entity.Employee) error {
	err := e.repo.Create(employee)
	if err != nil {
		return err
	}

	return nil
}
