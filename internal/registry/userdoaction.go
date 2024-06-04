package registry

import (
	"go-server/internal/api/handler"
	"go-server/internal/infrastructure/repository"
	"go-server/internal/usecase/userdoaction"
)

func (i *interactor) NewActionRepository() *repository.ActionRepo {
	return repository.NewActionRepo(i.mongo)
}

func (i *interactor) NewTransactionRepository() *repository.TransactionRepo {
	return repository.NewTransactionRepo(i.mongo)
}

func (i *interactor) NewUserDoActionService() *userdoaction.Service {
	return userdoaction.NewService(i.NewActionRepository(), i.NewAttendanceRepository(), i.NewTransactionRepository())
}

func (i *interactor) NewUserDoActionHandler() handler.UserDoActionHandler {
	return handler.NewUserDoActionHandler(i.NewUserDoActionService())
}
