package repository

import (
	"context"
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var TransactionCollectionName = "transactions"

type TransactionRepo struct {
	dbMongo *mongo.Database
}

func NewTransactionRepo(mg *mongo.Database) *TransactionRepo {
	return &TransactionRepo{dbMongo: mg}
}

func (repo *TransactionRepo) InsertTransaction(ctx context.Context, target *entity.Transaction) (*entity.Transaction, error) {
	result, err := repo.dbMongo.Collection(TransactionCollectionName).InsertOne(ctx, target)
	if err != nil {
		return nil, err
	}

	target.ID = commonentity.ID(result.InsertedID.(primitive.ObjectID).Hex())

	return target, nil
}
