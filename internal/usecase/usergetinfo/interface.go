package usergetinfo

import (
	"context"
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/entity"
)

type Attendance interface {
	SelectAttendanceList(ctx context.Context, userID commonentity.ID) (*[]entity.Attendance, error)
}

type Repository interface {
	Attendance
}

type UseCase interface {
	UserGetAttendanceList(ctx context.Context, userID commonentity.ID) (*[]entity.Attendance, error)
}
