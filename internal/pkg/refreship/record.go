package refreship

import "github.com/google/uuid"

type Status uint8

const (
	notStarted Status = iota
	getIpMgr
	getIp
	getItfMgr
	getItf
	updateItf
	succeed
	fail
)

type Record struct {
	ID      uuid.UUID `json:"id"`
	Status  Status    `json:"status"`
	Message string    `json:"message"`
}
