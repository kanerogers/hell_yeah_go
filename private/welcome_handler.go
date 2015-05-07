package private

import (
	"fmt"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/kanerogers/hell_yeah_go/common"
)

type WelcomeHandler struct {
}

type User struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (*WelcomeHandler) GetStatus(name string, user_table Table) common.Status {
	// Get user
	user_key := &dynamodb.Key{HashKey: "abc123"}
	user := new(User)
	err := user_table.GetDocument(user_key, user)

	if err != nil {
		fmt.Printf("Oh shit, there was an error: %s", err)
		panic(err)
	}

	first_name := user.FirstName
	greeting := fmt.Sprintf("Welcome to private town, %s", first_name)

	return common.Status{Status: greeting}
}
