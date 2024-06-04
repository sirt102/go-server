package entity

import commonentity "go-server/internal/common/cmentity"

type Action struct {
	commonentity.Entity `bson:",inline" json:",inline"`
	Type                string          `bson:"type" json:"type" binding:"omitempty,oneof=checkin checkout"`
	PlaceID             commonentity.ID `bson:"place_id" json:"place_id" binding:"required"`
	EmployeeID          commonentity.ID `bson:"employee_id" json:"employee_id"`
	Memo                string          `bson:"memo" json:"memo"`
}

type ActionType string

const (
	ActionCheckIn  ActionType = "checkin"
	ActionCheckOut ActionType = "checkout"
)
