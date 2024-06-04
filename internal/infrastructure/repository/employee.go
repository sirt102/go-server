package repository

import (
	"context"
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var EmployeeCollectionName = "employees"

type EmployeeRepo struct {
	dbMongo *mongo.Database
}

func NewEmployeeRepo(mg *mongo.Database) *EmployeeRepo {
	return &EmployeeRepo{dbMongo: mg}
}

func (repo *EmployeeRepo) InsertOne(ctx context.Context, target *entity.Employee) (*entity.Employee, error) {
	result, err := repo.dbMongo.Collection(EmployeeCollectionName).InsertOne(ctx, target)
	if err != nil {
		return nil, err
	}

	target.ID = commonentity.ID(result.InsertedID.(primitive.ObjectID).Hex())

	return target, nil
}
