package response

import (
	"encoding/json"

	"github.com/Olapat/go-architecture/model/entity"

	"github.com/gofiber/fiber/v2"
)

func ResponseOK(c *fiber.Ctx, status int, data interface{}, message string) error {
	isSuccess := status >= 200 && status < 300
	_message := ""
	if message != "" {
		_message = message
	} else if isSuccess {
		_message = "Success"
	} else if !isSuccess {
		_message = "Someting went wrong"
	}

	res := entity.Response{
		Success: isSuccess,
		Data:    data,
		Message: _message,
	}
	c.Set("Content-type", "application/json; charset=utf-8")
	j, _ := json.Marshal(res)
	return c.Status(status).Send(j)
}

func ResponseError(c *fiber.Ctx, status int, message string, errorRes interface{}) error {
	_message := ""
	if message != "" {
		_message = message
	} else if status == fiber.StatusBadRequest {
		_message = fiber.ErrBadRequest.Message
	} else if status == fiber.StatusNotFound {
		_message = fiber.ErrNotFound.Message
	} else if status == fiber.StatusUnauthorized {
		_message = fiber.ErrUnauthorized.Message
	} else {
		_message = fiber.ErrInternalServerError.Message
	}

	res := entity.ResponseError{
		Success: false,
		Data:    nil,
		Message: _message,
		Error:   errorRes,
	}
	c.Set("Content-type", "application/json; charset=utf-8")
	j, _ := json.Marshal(res)
	return c.Status(status).Send(j)
}
