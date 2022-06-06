package handler

import (
	"net/http"

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

	var o objectvalue.SignUpAccount

	if err := c.BodyParser(&o); err != nil {
		r := &response.SignUpAccountResponse{
			CommonResponse: response.CommonResponse{
				Status:  false,
				Code:    http.StatusUnprocessableEntity,
				Message: err.Error(),
			},
			Account: nil,
		}
		return c.Status(400).JSON(r)
	}

	if err := o.Validate(); err != nil {
		r := &response.SignUpAccountResponse{
			CommonResponse: response.CommonResponse{
				Status:  false,
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
			Account: nil,
		}
		return c.Status(400).JSON(r)

	}
	ctx := c.Context()
	r := a.accountUsecase.SignUp(ctx, &o)

	return c.Status(r.CommonResponse.Code).JSON(r)
}
