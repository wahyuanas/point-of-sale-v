package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyuanas/point-of-sale-v/account"
	objectvalue "github.com/wahyuanas/point-of-sale-v/account/api/object-value"
	"github.com/wahyuanas/point-of-sale-v/account/api/response"
)

type AccountHandler struct {
	accountUsecase account.AccountUsecase
}

func NewAccountHandler(app *fiber.App, a account.AccountUsecase) {
	handler := &AccountHandler{
		accountUsecase: a,
	}
	app.Post("/account", handler.SignUp)
}

func (a *AccountHandler) SignUp(c *fiber.Ctx) error {
	var o objectvalue.SignUp

	if err := c.BodyParser(&o); err != nil {
		r := &response.SignUpAccountResponse{
			CommonResponse: response.CommonResponse{
				Status:  false,
				Code:    fiber.StatusUnprocessableEntity,
				Message: err.Error(),
			},
		}
		r.Data.Account = nil
		return c.Status(fiber.StatusUnprocessableEntity).JSON(r)
	}

	if err := o.Validate(); err != nil {
		r := &response.SignUpAccountResponse{
			CommonResponse: response.CommonResponse{
				Status:  false,
				Code:    fiber.StatusBadRequest,
				Message: err.Error(),
			},
		}
		r.Data.Account = nil
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	ctx := c.Context()
	r := a.accountUsecase.SignUp(ctx, &o)

	return c.Status(r.CommonResponse.Code).JSON(r)
}
