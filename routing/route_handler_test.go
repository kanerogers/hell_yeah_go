package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/kanerogers/hell_yeah_go/common"
	checking "github.com/kanerogers/hell_yeah_go/testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type MockTableMaker struct {
}

func (m *MockTableMaker) GetTable(resourceName string) common.Table {
	table := new(MockTable)
	return table
}

type MockTable struct {
}

func (m *MockTable) GetDocument(key *dynamodb.Key, v interface{}) error {
	resourceName := key.HashKey
	resource := reflect.ValueOf(v).Elem()
	resource.FieldByName("ID").SetString(resourceName)
	return nil
}

func (m *MockTable) PutDocument(key *dynamodb.Key, v interface{}) error {
	resource := reflect.ValueOf(v).Elem()
	resource.FieldByName("ID").SetString("testing")

	return nil
}

func (m *MockTable) DeleteDocument(key *dynamodb.Key) error {
	return nil
}

type TestRoute struct {
	Method       string
	ExpectedCode int
	Path         string
	Data         interface{}
}

func TestRoutes(t *testing.T) {
	testRoutes := []TestRoute{
		{"GET", 200, "/api/listing/1", &common.Listing{ID: "1"}},
		{"POST", 201, "/api/listing", &common.Listing{ID: "testing"}},
		{"PUT", 200, "/api/listing/testing", &common.Listing{ID: "testing", Name: "Testing"}},
		{"DELETE", 204, "/api/listing/1", ""},
	}

	for _, testRoute := range testRoutes {
		actualCode, actualBody := testListing(testRoute.Method, testRoute.Path, testRoute.Data)

		assert.Equal(t, testRoute.ExpectedCode, actualCode)
		assert.Equal(t, testRoute.Data, actualBody)
	}

}

func testListing(testMethod string, testPath string, data interface{}) (responseCode int, response interface{}) {
	tableMaker := new(MockTableMaker)
	core := Core(tableMaker)

	testJSON, _ := json.Marshal(data)
	testJSONBuffer := bytes.NewBuffer(testJSON)

	fmt.Printf("About to send %s as our buffer\n", testJSONBuffer)

	r := checking.PerformRequest(core, testMethod, testPath, testJSONBuffer)

	actualCode := r.Code

	// If it's a DELETE, we want to send back an empty body, so our "response" is "nil"
	if testMethod == "DELETE" {
		return actualCode, r.Body.String()
	}

	// In all other cases, decode the Body as JSON into a Struct, then return it.
	actualBody := new(common.Listing)
	if err := json.Unmarshal(r.Body.Bytes(), actualBody); err != nil {
		fmt.Printf("Error unmarshaling JSON for some reason: %e\n", err)
		panic(err)
	}

	return actualCode, actualBody

}
