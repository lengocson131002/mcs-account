package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lengocson131002/mcs-account/domain/account"
)

type AccountController struct {
}

func NewAccountController() *AccountController {
	return &AccountController{}
}

// Swagger information here
func (c *AccountController) GetAccountBalance(ctx *fiber.Ctx) error {
	return RequestHandler[*account.CheckBalanceRequest, *account.CheckBalanceResponse](ctx)
}
