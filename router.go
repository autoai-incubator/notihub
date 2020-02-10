package main

import (
	"github.com/autoai-incubator/notihub/components"
	"github.com/gin-gonic/gin"
)

// NewRouter returns a new router
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
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
	r.POST("/:service/send/:recipient", func(c *gin.Context) {
		service := c.Param("service")
		var sr SendRequest
		c.BindJSON(&sr)
		if service == "telegram" {
			b := components.NewTelegramBot()
			user := components.GetUser(c.Param("recipient"))
			b.Send(&user, sr.Msg)
		}
	})
	return r
}
