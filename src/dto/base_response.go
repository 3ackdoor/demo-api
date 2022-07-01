package dto

import "github.com/3ackdoor/go-demo-api/src/type/null"

type BaseResponse struct {
	ID        uint        `json:"id"`
	CreatedAt null.Time   `json:"createdAt"`
	CreatedBy null.String `json:"createdBy"`
	UpdatedAt null.Time   `json:"updatedAt"`
	UpdatedBy null.String `json:"updatedBy"`
	DeletedAt null.Time   `json:"deletedAt"`
	DeletedBy null.String `json:"deletedBy"`
}
