package account

import (
	"context"

	objectvalue "github.com/wahyuanas/point-of-sale-v/account/api/object-value"
	"github.com/wahyuanas/point-of-sale-v/account/entity"
)

type AccountRepository interface {
	Store(ctx context.Context, cmd *objectvalue.SignUpAccount) (*entity.Account, interface{}, error)
}
