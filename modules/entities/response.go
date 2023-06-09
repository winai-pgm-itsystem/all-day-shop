package entities

import (
	"github.com/gofiber/fiber/v2"
	"github.com/winai-pgm-itsystem/all-day-shop/pkg/alldaylogger"
)

type IResponse interface {
	Success(code int, datat any) IResponse
	Error(code int, tractId, msg string) IResponse
	Res() error
}

type Response struct {
	StatusCode int
	Data       any
	ErrorRes   *ErrorResponse
	Context    *fiber.Ctx
	IsError    bool
}

type ErrorResponse struct {
	TraceId string `json:"trace_id"`
	Msg     string `json:"message"`
}

func NewResponse(c *fiber.Ctx) IResponse {
	return &Response{
		Context: c,
	}
}

func (r *Response) Success(code int, data any) IResponse {
	r.StatusCode = code
	r.Data = data
	alldaylogger.InitAlldayLogger(r.Context, code, &r.Data).Print().Save()
	return r
}
func (r *Response) Error(code int, tractId, msg string) IResponse {
	r.StatusCode = code
	r.ErrorRes = &ErrorResponse{
		TraceId: tractId,
		Msg:     msg,
	}
	r.IsError = true
	alldaylogger.InitAlldayLogger(r.Context, code, &r.ErrorRes).Print().Save()
	return r
}
func (r *Response) Res() error {
	return r.Context.Status(r.StatusCode).JSON(func() any {
		if r.IsError {
			return &r.ErrorRes

		}
		return &r.Data
	}())

}
