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

func (au *AccountUsecase) SignUp(ctx context.Context, cmd *objectvalue.SignUp) *response.SignUpAccountResponse {

	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	a, i, err := au.accountRepo.Store(ctx, cmd)
	if err != nil {
		r := &response.SignUpAccountResponse{
			CommonResponse: response.CommonResponse{
				Status:  false,
				Code:    i.(int),
				Message: err.Error(),
			},
		}
		r.Data.Account = nil

		return r
	}
	r := &response.SignUpAccountResponse{
		CommonResponse: response.CommonResponse{
			Status:  true,
			Code:    i.(int),
			Message: "Success",
		}}
	r.Data.Account = a

	return r

}
