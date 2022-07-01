package dto

import "github.com/3ackdoor/go-demo-api/src/type/null"

type BaseResponse struct {
	ID        uint      `json:"id"`
	CreatedAt null.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt null.Time `json:"updatedAt"`
	UpdatedBy string    `json:"updatedBy"`
	DeletedAt null.Time `json:"deletedAt"`
	DeletedBy string    `json:"deletedBy"`
}
