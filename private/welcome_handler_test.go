package private

import (
	"fmt"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/kanerogers/hell_yeah_go/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type MockTable struct {
	mock.Mock
}

func (m *MockTable) GetDocument(key *dynamodb.Key, v interface{}) error {
	user := reflect.ValueOf(v).Elem()

	user.Field(0).SetString("abc123")
	user.Field(1).SetString("Bede")
	user.Field(2).SetString("Overend")

	args := m.Called(key, v)
	return args.Error(0)
}

func TestWelcomeHandler(t *testing.T) {
	// Setup COT
	testHandler := new(WelcomeHandler)
	testName := "Bede"

	expectedStatusString := fmt.Sprintf("Welcome to private town, %s", testName)
	expectedStatus := hell_yeah_go.Status{Status: expectedStatusString}

	user_key := &dynamodb.Key{HashKey: "abc123"}
	user := User{"abc123", "Bede", "Overend"}

	mockTable := new(MockTable)
	mockTable.On("GetDocument", user_key, &user).Return(nil)

	// Excersising COT
	testStatus := testHandler.GetStatus(testName, mockTable)

	// Asserting
	assert.Equal(t, expectedStatus, testStatus)

	// (Optional) Cleanup
}
