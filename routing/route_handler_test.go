package routing

import (
	"encoding/json"
	"github.com/AdRoll/goamz/dynamodb"
	"github.com/kanerogers/hell_yeah_go/private"
	checking "github.com/kanerogers/hell_yeah_go/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"testing"
)

type MockTableMaker struct {
	mock.Mock
}

func (m *MockTableMaker) GetTable(resourceName string) private.Table {

	table := new(MockTable)
	m.Called(resourceName)

	return table
}

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

type Test struct {
	Status string `json:"status"`
}

func TestTakeListing(t *testing.T) {

	testPath := "/api/listing/1"
	testMethod := "GET"

	expectedCode := http.StatusOK
	expectedBody := &Test{Status: "take"}

	// Charset is because gin is funky like that
	expectedContentType := "application/json; charset=utf-8"

	tableMaker := new(MockTableMaker)
	core := Core(tableMaker)

	r := checking.PerformRequest(core, testMethod, testPath)

	actualCode := r.Code
	actualContentType := r.Header().Get("Content-Type")

	actualBody := new(Test)
	if err := json.Unmarshal(r.Body.Bytes(), actualBody); err != nil {
		panic(err)
	}

	assert.Equal(t, expectedCode, actualCode)
	assert.Equal(t, expectedContentType, actualContentType)
	assert.Equal(t, expectedBody, actualBody)

}
