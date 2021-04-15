package controllers

import (
	"fmt"
	"net/http"
	"transfer-api/bank/database"
	"transfer-api/bank/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	TransferController struct{}

	CreateTransferInput struct {
		AccountOriginID      string `json:"account_origin_id" validate:"required,uuid4"`
		AccountDestinationID string `json:"account_destination_id" validate:"required,uuid4"`
		Amount               int64  `json:"amount" validate:"gt=0,required"`
	}
)

func (tc *TransferController) InitiateTransfer(c echo.Context) (err error) {
	t := new(CreateTransferInput)
	if err = c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	acc := models.FindAccountByID(t.AccountOriginID)
	if acc.ID != "" {
		if acc.Balance < models.Money(t.Amount) {
			return c.JSON(http.StatusOK, "Insufficient fund")
		}

		acc.Balance -= models.Money(t.Amount)
	}

	accD := models.FindAccountByID(t.AccountDestinationID)
	accD.Balance += models.Money(t.Amount)

	models.UpdateAccountBalance(string(accD.ID), accD.Balance)
	models.UpdateAccountBalance(string(acc.ID), acc.Balance)

	database.DBConn.Transaction(func(tx *gorm.DB) error {
		return nil
	})

	return c.JSON(http.StatusOK, "Transaction successful")
}

func process(input *CreateTransferInput) error {
	acc := models.FindAccountByID("5ea25ebe-5222-4ed1-9acf-6ccdccbc35e0")
	if acc.ID != "" {

	}
	fmt.Println(acc)
	// origin, err := t.accountRepo.FindByID(ctx, domain.AccountID(input.AccountOriginID))
	// if err != nil {
	// 	switch err {
	// 	case domain.ErrAccountNotFound:
	// 		return domain.ErrAccountOriginNotFound
	// 	default:
	// 		return err
	// 	}
	// }

	// if err := origin.Withdraw(domain.Money(input.Amount)); err != nil {
	// 	return err
	// }

	// destination, err := t.accountRepo.FindByID(ctx, domain.AccountID(input.AccountDestinationID))
	// if err != nil {
	// 	switch err {
	// 	case domain.ErrAccountNotFound:
	// 		return domain.ErrAccountDestinationNotFound
	// 	default:
	// 		return err
	// 	}
	// }

	// destination.Deposit(domain.Money(input.Amount))

	// if err = t.accountRepo.UpdateBalance(ctx, origin.ID(), origin.Balance()); err != nil {
	// 	return err
	// }

	// if err = t.accountRepo.UpdateBalance(ctx, destination.ID(), destination.Balance()); err != nil {
	// 	return err
	// }

	return nil
}
