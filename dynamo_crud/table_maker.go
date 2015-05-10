package dynamo_crud

import (
	"fmt"
	"github.com/AdRoll/goamz/aws"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/kanerogers/hell_yeah_go/common"
)

type TableMaker interface {
	GetTable(string) common.Table
}

type DynamoTableMaker struct {
}

func (t *DynamoTableMaker) GetTable(resourceName string) common.Table {

	primaryKeyName := fmt.Sprintf("%s_id", resourceName)

	// AWS Authentication.
	auth, err := aws.EnvAuth()

	if err != nil {
		errorMessage := fmt.Sprintf("Authentication error: %s", err)
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
