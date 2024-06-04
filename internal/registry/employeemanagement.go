package registry

import (
	"go-server/internal/api/handler"
	"go-server/internal/infrastructure/repository"
	"go-server/internal/usecase/employeemanagement"
)

func (i *interactor) NewEmployeeManagementRepository() *repository.EmployeeRepo {
	return repository.NewEmployeeRepo(i.mongo)
}

func (i *interactor) NewEmployeeService() *employeemanagement.Service {
	return employeemanagement.NewService(i.NewEmployeeManagementRepository())
}

func (i *interactor) NewEmployeeManagementHandler() handler.EmployeemManagementHandler {
	return handler.NewEmployeeManagementHandler(i.NewEmployeeService())
}
