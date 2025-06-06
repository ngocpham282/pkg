package apif

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg/errs"
)

type Response struct {
	Metadata Metadata `json:"metadata"`
	Data     any      `json:"data,omitempty"`
}

type Metadata struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Details any    `json:"details,omitempty"`
}

func Respond(ctx *gin.Context, data any, err error) {
	httpCode := http.StatusOK
	response := Response{
		Metadata: Metadata{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data: data,
	}

	if err != nil {
		var appErr *errs.AppError
		switch {
		case errors.As(err, &appErr):
			httpCode = appErr.Code / 1000
			response.Metadata.Message = appErr.Message
			response.Metadata.Code = appErr.Code
			response.Metadata.Details = appErr.Details
		default:
			httpCode = http.StatusInternalServerError
			response.Metadata.Code = http.StatusInternalServerError
			response.Metadata.Message = "Internal Server Error"
			response.Metadata.Details = err.Error()
		}
	}

	ctx.JSON(httpCode, response)
}
