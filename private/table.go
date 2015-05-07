package private

import (
	"github.com/AdRoll/goamz/aws"
	"github.com/AdRoll/goamz/dynamodb"
)

type Table interface {
	GetDocument(key *dynamodb.Key, v interface{}) error
}

func getTable() dynamodb.Table {
	// AWS Authentication.
	auth := aws.Auth{AccessKey: "derp", SecretKey: "dope"}
	region := aws.GetRegion("ap-southeast-2")

	// Setup a server.
	server := dynamodb.New(auth, region)

	// Setup table.
	primary_key_attribute := &dynamodb.Attribute{Type: "S", Name: "user_id"}
	primary_key := dynamodb.PrimaryKey{KeyAttribute: primary_key_attribute}
	table := dynamodb.Table{Server: server, Name: "sproutli_users", Key: primary_key}

	return table
}
