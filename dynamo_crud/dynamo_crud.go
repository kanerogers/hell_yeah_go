package dynamo_crud

import (
	"fmt"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/kanerogers/hell_yeah_go/private"
	"net/http"
)

type DynamoCRUD struct {
	private.Table
	Resource interface{}
}

func NewDynomoCRUD(table private.Table, resourceName string, data interface{}) DynamoCRUD {
	return DynamoCRUD{
		table,
		data,
	}
}

func (dynamo *DynamoCRUD) CreateHandler(c *gin.Context) {
	fmt.Println("Creating")
	c.JSON(http.StatusOK, gin.H{"status": "create"})
}

func (dynamo *DynamoCRUD) TakeHandler(c *gin.Context) {
	resourceName := c.Params[0].Value
	fmt.Printf("Taking %s\n", resourceName)
	key := &dynamodb.Key{HashKey: resourceName}

	err := dynamo.Table.GetDocument(key, dynamo.Resource)

	if err != nil {
		c.Fail(500, err)
		return
	}

	fmt.Printf("Found %s\n", dynamo.Resource)

	c.JSON(http.StatusOK, dynamo.Resource)
}
