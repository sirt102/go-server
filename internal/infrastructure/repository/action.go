package repository

import (
	"context"
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ActionCollectionName = "actions"

type ActionRepo struct {
	dbMongo *mongo.Database
}

func NewActionRepo(mg *mongo.Database) *ActionRepo {
	return &ActionRepo{dbMongo: mg}
}

func (repo *ActionRepo) InsertAction(ctx context.Context, target *entity.Action) (*entity.Action, error) {
	result, err := repo.dbMongo.Collection(ActionCollectionName).InsertOne(ctx, target)
	if err != nil {
		return nil, err
	}

	target.ID = commonentity.ID(result.InsertedID.(primitive.ObjectID).Hex())

	return target, nil
}
