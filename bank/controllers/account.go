package controllers

import (
	"fmt"
	"net/http"
	"time"
	"transfer-api/bank/common"
	"transfer-api/bank/models"

	"github.com/labstack/echo/v4"
)

type (
	CreateAccountInput struct {
		Name    string `json:"name" validate:"required"`
		Balance int64  `json:"balance" validate:"gt=0,required"`
	}

	CreateAccountPresenter interface {
		Output(models.Account) CreateAccountOutput
	}

	CreateAccountOutput struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		Balance   float64 `json:"balance"`
		CreatedAt string  `json:"created_at"`
	}

	AccountController struct {
		repo      models.AccountRepository
		presenter CreateAccountPresenter
	}
)

func (tc *AccountController) Create(c echo.Context) (err error) {
	a := new(models.Account)
	if err = c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var account = models.NewAccount(
		models.AccountID(common.NewUUID()),
		a.Name,
		models.Money(a.Balance),
		time.Now(),
	)

	fmt.Println(account)
	models.CreateAccountModel(account)
	return c.JSON(http.StatusOK, Output(account))
}

// presenter
func Output(account models.Account) CreateAccountOutput {
	return CreateAccountOutput{
		ID:   string(account.ID),
		Name: string(account.Name),
	}
}
