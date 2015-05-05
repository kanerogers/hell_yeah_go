package main

import (
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/kanerogers/test_app/common"
	"github.com/kanerogers/test_app/private"
)

func main() {
	router := gin.Default()
	welcome_handler := new(private.WelcomeHandler)

	private_router := router.Group("/api/private")
	private_router.Use(jwt.Auth(test_app.JWTKEY))

	private_router.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, welcome_handler.GetStatus())
	})

	router.Run(":8000")
}
