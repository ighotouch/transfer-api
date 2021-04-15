package presenter

import (
	"transfer-api/bank/controllers"
	"transfer-api/bank/models"
)

type createAccountPresenter struct{}

func NewCreateAccountPresenter() controllers.CreateAccountPresenter {
	return createAccountPresenter{}
}

func (a createAccountPresenter) Output(account models.Account) controllers.CreateAccountOutput {
	return controllers.CreateAccountOutput{
		ID:   string(account.ID),
		Name: string(account.Name),
	}
}
