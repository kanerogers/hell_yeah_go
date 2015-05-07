package main

import (
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/contrib/rest"
	"github.com/gin-gonic/gin"
	"github.com/kanerogers/hell_yeah_go/common"
	"github.com/kanerogers/hell_yeah_go/listing"
	"github.com/kanerogers/hell_yeah_go/private"
)

func main() {
	router := gin.Default()
	// welcome_handler := new(private.WelcomeHandler)

	private_router := router.Group("/api/private")
	private_router.Use(jwt.Auth(common.JWTKEY))

	// private_router.GET("/welcome", func(c *gin.Context) {
	// 	c.JSON(200, welcome_handler.GetStatus("Bede"))
	// })

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"Hello": "World"})
	})

	api := router.Group("/api")
	rest.CRUD(api, "/listing", new(listing.Handler))

	router.Run(":8000")
}
