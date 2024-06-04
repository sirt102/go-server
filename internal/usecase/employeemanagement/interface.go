package employeemanagement

import (
	"context"
	"go-server/internal/entity"
)

type Action interface {
	InsertOne(ctx context.Context, ac *entity.Employee) (*entity.Employee, error)
}

type Repository interface {
	Action
}

type UseCase interface {
	CreateNewEmployee(ctx context.Context, em *entity.Employee) (*entity.Employee, error)
}
