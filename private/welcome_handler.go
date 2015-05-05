package private

import (
	"fmt"
	"github.com/AdRoll/goamz/aws"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/kanerogers/test_app/common"
)

type WelcomeHandler struct {
}

type User struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (*WelcomeHandler) GetStatus(name string) test_app.Status {
	// AWS Authentication.
	auth := aws.Auth{AccessKey: "AKIAIBL23CYNOCMPOJ7Q", SecretKey: "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"}
	region := aws.GetRegion("ap-southeast-2")

	// Setup a server.
	server := dynamodb.New(auth, region)

	// Setup table.
	primary_key_attribute := &dynamodb.Attribute{Type: "S", Name: "user_id"}
	primary_key := dynamodb.PrimaryKey{KeyAttribute: primary_key_attribute}
	table := dynamodb.Table{Server: server, Name: "sproutli_users", Key: primary_key}

	// Get user
	user_key := &dynamodb.Key{HashKey: "abc123"}
	user := new(User)
	err := table.GetDocument(user_key, user)

	if err != nil {
		fmt.Printf("Oh shit, there was an error: %s", err)
		panic(err)
	}

	first_name := user.FirstName
	greeting := fmt.Sprintf("Welcome to private town, %s", first_name)

	return test_app.Status{Status: greeting}
}
