package usergetinfo

import (
	"context"
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/entity"
	"log"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) UserGetAttendanceList(ctx context.Context, userID commonentity.ID) (attendanceList *[]entity.Attendance, err error) {
	attendanceList, err = s.repo.SelectAttendanceList(ctx, userID)
	if err != nil {
		// TODO: Send notify to our system
		log.Println("[UserGetAttendanceList] - [SelectAttendanceList] - ", err.Error())
	}

	return
}
