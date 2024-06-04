package entity

import (
	commonentity "go-server/internal/common/cmentity"
	"time"
)

type Attendance struct {
	commonentity.Entity `bson:",inline" json:",inline"`
	CheckInTime         time.Time       `bson:"check_in_time" json:"check_in_time"`
	CheckOutTime        time.Time       `bson:"check_out_time" json:"check_out_time"`
	Date                string          `bson:"date" json:"date"`
	OverTime            string          `bson:"over_time" json:"over_time"`
	TotalHour           string          `bson:"total_hour" json:"total_hour"`
	EmployeeID          commonentity.ID `bson:"employee_id" json:"employee_id" binding:"required"`
}

type AttendanceUpdate struct {
	Set AttendanceUpdateSet `bson:"$set"`
}

type AttendanceUpdateSet struct {
	CheckInTime  *time.Time      `bson:"check_in_time,omitempty"`
	CheckOutTime *time.Time      `bson:"check_out_time,omitempty"`
	OverTime     *string         `bson:"over_time,omitempty"`
	TotalHour    *string         `bson:"total_hour,omitempty"`
	EmployeeID   commonentity.ID `bson:"employee_id" binding:"required"`
}
