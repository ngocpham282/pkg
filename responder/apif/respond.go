package apif

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	response := Response{
		Metadata: Metadata{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data: data,
	}

	if err != nil {
		response.Metadata.Code = 500
		response.Metadata.Message = "Internal Server Error"
		response.Metadata.Details = err.Error()
	}

	ctx.JSON(response.Metadata.Code, response)
}
