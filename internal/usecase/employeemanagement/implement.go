package employeemanagement

import (
	"context"
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

func (s *Service) CreateNewEmployee(ctx context.Context, em *entity.Employee) (*entity.Employee, error) {
	insertedEmploee, err := s.repo.InsertOne(ctx, em)
	if err != nil {
		log.Println("[CreateNewEmployee] - [InsertOne] - ", err.Error())
		return nil, err
	}

	return insertedEmploee, nil
}
