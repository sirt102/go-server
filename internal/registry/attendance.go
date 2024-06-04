package registry

import (
	"go-server/internal/api/handler"
	"go-server/internal/infrastructure/repository"
	"go-server/internal/usecase/usergetinfo"
)

func (i *interactor) NewUserGetInfoService() *usergetinfo.Service {
	return usergetinfo.NewService(i.NewAttendanceRepository())
}

func (i *interactor) NewAttendanceRepository() *repository.AttendanceRepo {
	return repository.NewAttendanceRepo(i.mongo)
}

func (i *interactor) NewUserGetInfoHandler() handler.UserGetInfoHandler {
	return handler.NewUserGetInfoHandler(i.NewUserGetInfoService())
}
