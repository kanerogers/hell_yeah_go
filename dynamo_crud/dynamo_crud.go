package dynamo_crud

import (
	"fmt"
	"github.com/AdRoll/goamz/aws"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/kanerogers/hell_yeah_go/private"
	"net/http"
)

type DynamoCRUD struct {
	private.Table
	Resource interface{}
}

func NewDynomoCRUD(resourceName string, data interface{}) DynamoCRUD {
	return DynamoCRUD{
		getTable(resourceName),
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

func getTable(resourceName string) private.Table {
	primaryKeyName := fmt.Sprintf("%s_id", resourceName)

	// AWS Authentication.
	auth, err := aws.EnvAuth()

	if err != nil {
		errorMessage := fmt.Sprintf("Authentication error: %e", err)
		panic(errorMessage)
	}

	region := aws.GetRegion("ap-southeast-2")

	// Setup a server.
	server := dynamodb.New(auth, region)

	// Setup table.
	primary_key_attribute := &dynamodb.Attribute{Type: "S", Name: primaryKeyName}
	primary_key := dynamodb.PrimaryKey{KeyAttribute: primary_key_attribute}
	table := dynamodb.Table{Server: server, Name: resourceName, Key: primary_key}

	return &table
}
