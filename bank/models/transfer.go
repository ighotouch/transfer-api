package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type TransferID string

type Money int64

func (t TransferID) String() string {
	return string(t)
}

type (
	TransferRepository interface {
		Create(context.Context, Transfer) (Transfer, error)
		FindAll(context.Context) ([]Transfer, error)
		WithTransaction(context.Context, func(context.Context) error) error
	}

	Transfer struct {
		gorm.Model
		Id                   TransferID
		AccountOriginID      AccountID `json:"AccountID" validate:"required"`
		AccountDestinationID AccountID `json:"AccountDestinationID" validate:"required"`
		Amount               Money     `json:"Amount" validate:"required"`
		CreatedAt            time.Time
	}
)
