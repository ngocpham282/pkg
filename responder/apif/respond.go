package apif



type Response struct {
	Metadata Metadata `json:"metadata"`
	Data     any      `json:"data,omitempty"`
}

type Metadata struct {
	Code   int    `json:"code"`
	Message string `json:"message,omitempty"`
	Details any `json:"details,omitempty"`
}

func Respond(ctx *gin.Context, data any, err error) {
	logger := logging.FromContext(ctx)
	response := Response{
		Metadata: Metadata{
			Code:    200,
			Message: "OK",
		},
		Data: data,
	}

	if err != nil {
		logger.Errorw("Error in response", "error", err)
		response.Metadata.Code = 500
		response.Metadata.Message = "Internal Server Error"
		response.Metadata.Details = err.Error()
	}

	ctx.JSON(response.Metadata.Code, response)
}