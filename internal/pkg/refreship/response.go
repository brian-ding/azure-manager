package refreship

import "github.com/google/uuid"

type RefreshResponse struct {
	RecordId uuid.UUID `json:"recordId"`
}

type CheckResponse struct {
	Message string `json:"message"`
	Record  Record `json:"record"`
}
