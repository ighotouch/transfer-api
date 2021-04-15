package models

import (
	"context"
	"errors"
	"time"
	"transfer-api/bank/database"
)

var (
	ErrAccountNotFound = errors.New("account not found")

	ErrAccountOriginNotFound = errors.New("account origin not found")

	ErrAccountDestinationNotFound = errors.New("account destination not found")

	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")
)

type AccountID string

func (a AccountID) String() string {
	return string(a)
}

type (
	AccountRepository interface {
		Create(Account) (Account, error)
		UpdateBalance(context.Context, AccountID, Money) error
		FindAll(context.Context) ([]Account, error)
		FindByID(context.Context, AccountID) (Account, error)
		FindBalance(context.Context, AccountID) (Account, error)
	}

	Account struct {
		ID        AccountID
		Name      string `json:"name" validate:"required"`
		Balance   Money  `json:"balance"`
		CreatedAt time.Time
	}
)

var account Account

func NewAccount(ID AccountID, name string, balance Money, createdAt time.Time) Account {
	return Account{
		ID:        ID,
		Name:      name,
		Balance:   balance,
		CreatedAt: createdAt,
	}
}

func (a *Account) Deposit(amount Money) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount Money) error {
	if a.Balance < amount {
		return ErrInsufficientBalance
	}

	a.Balance -= amount

	return nil
}

func CreateAccountModel(acc Account) {
	database.DBConn.Create(acc)
}

func FindAccountByID(id string) Account {
	var acc []Account
	err := database.DBConn.Raw(`SELECT * FROM accounts WHERE id = ?`, id).Scan(&acc)

	if err != nil {
	}
	if len(acc) == 0 {
		return Account{}
	}

	return acc[0]
}

func UpdateAccountBalance(id string, balance Money) {
	database.DBConn.Exec(`UPDATE accounts SET balance = ? WHERE id = ?`, balance, id)
}
