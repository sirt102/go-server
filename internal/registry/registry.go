package registry

import (
	"go-server/internal/api/handler"
	"go-server/pkg/mongo"
)

type interactor struct {
	mongo mongo.MongoDB
}

// Interactor Interactor interface
type Interactor interface {
	NewAppHandler() handler.AppHandler
}

// NewInteractor Constructs new interactor
func NewInteractor(mg mongo.MongoDB) Interactor {
	return &interactor{mongo: mg}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return handler.AppHandler{
		EmployeemManagementHandler: i.NewEmployeeManagementHandler(),
		UserDoActionHandler:        i.NewUserDoActionHandler(),
		UserGetInfoHandler:         i.NewUserGetInfoHandler(),
	}
}
