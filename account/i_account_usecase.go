package account

import (
	"context"

	objectvalue "github.com/wahyuanas/point-of-sale-v/account/api/object-value"
	"github.com/wahyuanas/point-of-sale-v/account/api/response"
)

type AccountUsecase interface {
	SignUp(ctx context.Context, cmd *objectvalue.SignUpAccount) *response.SignUpAccountResponse
}
