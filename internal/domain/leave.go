package domain

import (
	"github.com/google/uuid"
)

type LeaveBalance struct {
	Base
	EmployeeID     uuid.UUID `json:"employee_id" gorm:"type:uuid;not null"`
	OrganizationID uuid.UUID `json:"organization_id" gorm:"type:uuid;not null"`
	LeaveTypeID    uuid.UUID `json:"leave_type_id" gorm:"type:uuid;not null"`
	Year           int       `json:"year" gorm:"not null"`
	TotalDays      float64   `json:"total_days" gorm:"type:numeric(5,2);not null"`
	UsedDays       float64   `json:"used_days" gorm:"type:numeric(5,2);default:0"`
	PendingDays    float64   `json:"pending_days" gorm:"type:numeric(5,2);default:0"`
}
