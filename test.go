package main

import (
	"fmt"
	// "github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/contrib/rest"
	"github.com/gin-gonic/gin"
	"github.com/kanerogers/hell_yeah_go/common"
	"github.com/kanerogers/hell_yeah_go/dynamo_crud"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"Hello": "World"})
	})

	api := router.Group("/api")
	// api.Use(jwt.Auth(common.JWTKEY))

	resources := map[string]interface{}{
		"listing": new(common.Listing),
		"user":    new(common.User),
	}

	for resourceName, resource := range resources {
		handler := dynamo_crud.NewDynomoCRUD(resourceName, resource)
		path := fmt.Sprintf("/%s", resourceName)
		rest.CRUD(api, path, &handler)
	}

	router.Run(":8000")
}
