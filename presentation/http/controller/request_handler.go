package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lengocson131002/go-clean-core/pipeline"
	"github.com/lengocson131002/go-clean-core/transport/http"
)

func RequestHandler[TReq any, TResp any](ctx *fiber.Ctx) error {
	// Step 1: Parse the request
	var req TReq
	err := ctx.BodyParser(req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	// Step 2: Send request to the request pipeline
	resp, err := pipeline.Send[TReq, TResp](ctx.UserContext(), req)
	if err != nil {
		return err
	}

	httpResp := http.SuccessResponse[TResp](resp)
	return ctx.Status(httpResp.Result.Status).JSON(httpResp)
}

func NotificationHandler[TNoti any](ctx *fiber.Ctx) error {
	// Step 1: Parse the request
	var req TNoti
	err := ctx.BodyParser(req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	// Step 2: Send notification to the request pipeline
	err = pipeline.Publish[TNoti](ctx.UserContext(), req)
	if err != nil {
		return err
	}

	httpResp := http.DefaultSuccessResponse
	return ctx.Status(httpResp.Result.Status).JSON(httpResp)
}
