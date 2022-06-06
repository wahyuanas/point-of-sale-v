package usecase

import (
	"context"
	"time"

	"github.com/wahyuanas/point-of-sale-v/account"
	objectvalue "github.com/wahyuanas/point-of-sale-v/account/api/object-value"
	"github.com/wahyuanas/point-of-sale-v/account/api/response"
)

type AccountUsecase struct {
	accountRepo    account.AccountRepository
	contextTimeout time.Duration
}

func NewAccountUsecase(a account.AccountRepository, t time.Duration) account.AccountUsecase {
	return &AccountUsecase{
		accountRepo:    a,
		contextTimeout: t,
	}
}

func (au *AccountUsecase) SignUp(ctx context.Context, cmd *objectvalue.SignUpAccount) *response.SignUpAccountResponse {

	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	a, i, err := au.accountRepo.Store(ctx, cmd)
	if err != nil {
		return &response.SignUpAccountResponse{
			CommonResponse: response.CommonResponse{
				Status:  false,
				Code:    i.(int),
				Message: err.Error(),
			},
			Account: nil,
		}
	}
	return &response.SignUpAccountResponse{
		CommonResponse: response.CommonResponse{
			Status:  true,
			Code:    i.(int),
			Message: "Success",
		},
		Account: a,
	}
}
