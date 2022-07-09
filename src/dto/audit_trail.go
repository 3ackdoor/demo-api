package dto

import (
	"github.com/3ackdoor/go-demo-api/src/type/null"
	"gorm.io/gorm"
)

type AuditTrail struct {
	CreatedAt null.Time   `json:"createdAt"`
	CreatedBy null.String `json:"createdBy"`
	UpdatedAt null.Time   `json:"updatedAt"`
	UpdatedBy null.String `json:"updatedBy"`
	DeletedAt gorm.DeletedAt   `json:"deletedAt"`
	DeletedBy null.String `json:"deletedBy"`
}
