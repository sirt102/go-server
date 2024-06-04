package userdoaction

import (
	"context"
	"go-server/internal/common/constant"
	"go-server/internal/entity"
	"go-server/pkg/blockchain/chain"
	"log"
	"time"
)

type Service struct {
	actionRepo      ActionRepo
	attendanceRepo  AttendanceRepo
	transactionRepo TransactionRepo
}

func NewService(actionRepo ActionRepo, attendanceRepo AttendanceRepo, transactionRepo TransactionRepo) *Service {
	return &Service{
		actionRepo:      actionRepo,
		attendanceRepo:  attendanceRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *Service) UserCreateAction(ctx context.Context, ac *entity.Action) (*entity.Action, error) {
	timeNow := time.Now()
	insertedAction, err := s.actionRepo.InsertAction(ctx, ac)
	if err != nil {
		// TODO: Notify to our system.
		log.Println("[UserCreateAction] - InsertOne] - ", err.Error())

		return nil, err
	}

	// Use switch case instead of if-else because in real, we have more than two actions: check-in and check-out
	switch ac.Type {
	case string(entity.ActionCheckIn):
		success, err := s.attendanceRepo.UpsertAttendance(ctx, ac.EmployeeID, &entity.AttendanceUpdate{
			Set: entity.AttendanceUpdateSet{
				CheckInTime: &timeNow,
				EmployeeID:  ac.EmployeeID,
			},
		})
		if err != nil {
			// TODO: Notify to our system.
			log.Println("[UserCreateAction] - [UpsertAttendance] - ", err.Error())
		}
		if !success {
			// TODO: Notify to our system.
			log.Println("[UserCreateAction] - [UpsertAttendance] - Update Attendance Fail!")
		}
	case string(entity.ActionCheckOut):
		currentAttendance, err := s.attendanceRepo.SelectAttendance(ctx, ac.EmployeeID)
		if err != nil {
			log.Println("[UserCreateAction] - [SelectAttendance] - ", err.Error())
		}
		var (
			totalTime                 = time.Since(currentAttendance.CheckInTime)
			totalHour                 = (totalTime - time.Duration(constant.REST_TIME_IN_NANO))
			totalHourString, overTime string
		)

		if totalHour > 0 {
			totalHourString = totalHour.String()
		}
		if totalHour > 8 {
			overTime = (totalHour - time.Duration(constant.STANDARD_WORKING_TIME)).String()
		}
		success, err := s.attendanceRepo.UpsertAttendance(ctx, ac.EmployeeID, &entity.AttendanceUpdate{
			Set: entity.AttendanceUpdateSet{
				CheckOutTime: &timeNow,
				TotalHour:    &totalHourString,
				OverTime:     &overTime,
				EmployeeID:   ac.EmployeeID,
			},
		})
		if err != nil {
			// TODO: Notify to our system.
			log.Println("[UserCreateAction] - [UpsertAttendance] - ", err.Error())
		}
		if !success {
			// TODO: Notify to our system.
			log.Println("[UserCreateAction] - [UpsertAttendance] - Update Attendance Fail!")
		}

		go func(ctx context.Context, data ...any) {
			web3Instance, err := chain.NewWeb3Instance()
			if err != nil {
				return
			}

			txHash, err := web3Instance.SendTransaction(ctx, data...)
			if err != nil {
				log.Println("Call to Smart Contract Error!")

				return
			}

			newTx, err := s.transactionRepo.InsertTransaction(ctx, &entity.Transaction{
				TransactionID: txHash,
			})
			if err != nil {
				// Notify to our system
				log.Println("[UserCreateAction] - [InsertTransaction] - ", err.Error())
			}

			log.Println("Call to Blockchain Successfully: ", newTx.TransactionID)
			//TODO : Add more process here to verify or do something else with new transaction
		}(ctx, currentAttendance.ID.String(), ac.EmployeeID.String(), currentAttendance.CheckInTime.String(), timeNow.String(), currentAttendance.Date)
	}

	return insertedAction, nil
}
