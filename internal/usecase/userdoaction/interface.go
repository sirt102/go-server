package userdoaction

import (
	"context"
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/entity"
)

type ActionRepo interface {
	InsertAction(ctx context.Context, ac *entity.Action) (*entity.Action, error)
}

type AttendanceRepo interface {
	SelectAttendance(ctx context.Context, userID commonentity.ID) (*entity.Attendance, error)
	UpsertAttendance(ctx context.Context, userID commonentity.ID, update *entity.AttendanceUpdate) (bool, error)
}

type TransactionRepo interface {
	InsertTransaction(ctx context.Context, target *entity.Transaction) (*entity.Transaction, error)
}

type UseCase interface {
	UserCreateAction(ctx context.Context, ac *entity.Action) (*entity.Action, error)
}
