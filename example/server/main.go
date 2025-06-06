package main

import (
	"github.com/gin-gonic/gin"
	"pkg/apif"
	"pkg/errs"
	"pkg/logging"
	"pkg/middleware"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.GET("/example", func(c *gin.Context) {
		logger := logging.From(c)
		logger.Info("Hello World")
		appErr := errs.NewAppError(errs.ErrBadRequest)
		apif.Respond(c, nil, appErr)
	})
	r.Run(":8080")
}
