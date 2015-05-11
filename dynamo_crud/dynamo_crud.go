package dynamo_crud

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/kanerogers/hell_yeah_go/common"
	"net/http"
	"reflect"
)

type DynamoCRUD struct {
	common.Table
	Resource interface{}
}

func NewDynomoCRUD(table common.Table, resourceName string, data interface{}) DynamoCRUD {
	return DynamoCRUD{
		table,
		data,
	}
}

func (dynamo *DynamoCRUD) CreateHandler(c *gin.Context) {
	// Marshal the POST from JSON to Struct
	c.Bind(dynamo.Resource)

	// Generate a random ID
	randomId := uuid.New()
	key := &dynamodb.Key{HashKey: randomId}

	// Set the ID on the Struct using reflection as we don't know the type here.)
	reflect.ValueOf(dynamo.Resource).Elem().FieldByName("ID").SetString(randomId)

	// Post the document to DynamoDB
	err := dynamo.Table.PutDocument(key, dynamo.Resource)

	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(201, dynamo.Resource)
}

func (dynamo *DynamoCRUD) TakeHandler(c *gin.Context) {
	resourceName := c.Params[0].Value
	key := &dynamodb.Key{HashKey: resourceName}

	err := dynamo.Table.GetDocument(key, dynamo.Resource)

	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(http.StatusOK, dynamo.Resource)
}

func (dynamo *DynamoCRUD) UpdateHandler(c *gin.Context) {
	// Marshal the PUT from JSON to Struct
	c.Bind(dynamo.Resource)

	fmt.Printf("Received %s via JSON PUT\n", c.Request.Body)

	resourceName := c.Params[0].Value
	key := &dynamodb.Key{HashKey: resourceName}

	err := dynamo.Table.PutDocument(key, dynamo.Resource)

	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(http.StatusOK, dynamo.Resource)
}

func (dynamo *DynamoCRUD) DeleteHandler(c *gin.Context) {
	resourceName := c.Params[0].Value
	key := &dynamodb.Key{HashKey: resourceName}

	err := dynamo.Table.DeleteDocument(key)

	if err != nil {
		c.Fail(500, err)
		return
	}

	c.Writer.WriteHeader(204)
}
