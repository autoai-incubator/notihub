package main

import (
	"github.com/autoai-incubator/notihub/components"
	"github.com/gin-gonic/gin"
)

// NewRouter returns a new router
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/:service/send", func(c *gin.Context) {
		service := c.Param("service")
		var sr SendRequest
		c.BindJSON(&sr)
		if service == "telegram" {
			b := components.NewTelegramBot()
			user := components.GetUser(sr.Recipient)
			b.Send(&user, sr.Msg)
		}
	})
	return r
}
