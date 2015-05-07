package routing

import (
	"fmt"
	"github.com/gin-gonic/contrib/rest"
	"github.com/gin-gonic/gin"
	"github.com/kanerogers/hell_yeah_go/common"
	"github.com/kanerogers/hell_yeah_go/dynamo_crud"
)

func Core(tableMaker dynamo_crud.TableMaker) *gin.Engine {

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
		table := tableMaker.GetTable(resourceName)
		handler := dynamo_crud.NewDynomoCRUD(table, resourceName, resource)
		path := fmt.Sprintf("/%s", resourceName)
		rest.CRUD(api, path, &handler)
	}

	return router
}
