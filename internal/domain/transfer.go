package domain

import (
	"github.com/google/uuid"
	"time"
)

type TransferStatus string

const (
	TransferPending   TransferStatus = "PENDING"
	TransferProcessed TransferStatus = "PROCESSED"
	TransferFailed    TransferStatus = "FAILED"
)

type Transfer struct {
	ID             uuid.UUID
	IdempotencyKey string
	FromWalletID   uuid.UUID
	ToWalletID     uuid.UUID
	Amount         int64
	Status         TransferStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
