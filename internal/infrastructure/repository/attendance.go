package repository

import (
	"context"
	commonentity "go-server/internal/common/cmentity"
	"go-server/internal/common/util"
	"go-server/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var AttendanceCollectionName = "attendances"

type AttendanceRepo struct {
	dbMongo *mongo.Database
}

func NewAttendanceRepo(mg *mongo.Database) *AttendanceRepo {
	return &AttendanceRepo{dbMongo: mg}
}

func (repo *AttendanceRepo) SelectAttendance(ctx context.Context, employeeID commonentity.ID) (*entity.Attendance, error) {
	var (
		filter = bson.D{
			{Key: "$and",
				Value: bson.A{
					bson.D{{Key: "date", Value: util.Today()}},
					bson.D{{Key: "employee_id", Value: employeeID}},
				},
			},
		}
		at entity.Attendance
	)

	err := repo.dbMongo.Collection(AttendanceCollectionName).FindOne(ctx, filter).Decode(&at)
	if err != nil {
		return nil, err
	}

	return &at, nil
}

func (repo *AttendanceRepo) UpsertAttendance(ctx context.Context, employeeID commonentity.ID, update *entity.AttendanceUpdate) (bool, error) {
	var (
		option = options.Update().SetUpsert(true)
		filter = bson.D{
			{Key: "$and",
				Value: bson.A{
					bson.D{{Key: "date", Value: util.Today()}},
					bson.D{{Key: "employee_id", Value: employeeID}},
				},
			},
		}
	)

	result, err := repo.dbMongo.Collection(AttendanceCollectionName).UpdateOne(
		ctx,
		filter,
		update,
		option,
	)
	if err != nil {
		return false, err
	}

	return result.UpsertedCount+result.ModifiedCount+result.MatchedCount > 0, nil
}

func (repo *AttendanceRepo) SelectAttendanceList(ctx context.Context, employeeID commonentity.ID) (*[]entity.Attendance, error) {
	cursor, err := repo.dbMongo.Collection(AttendanceCollectionName).Find(ctx, bson.M{"employee_id": employeeID})
	if err != nil {
		return nil, err
	}

	var attandanceList []entity.Attendance
	if err = cursor.All(ctx, &attandanceList); err != nil {
		return nil, err
	}

	return &attandanceList, nil
}
