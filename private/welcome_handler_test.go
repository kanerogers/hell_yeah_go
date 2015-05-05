package private

import (
	"fmt"
	"github.com/kanerogers/test_app/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWelcomeHandler(t *testing.T) {
	// Setup COT
	testHandler := new(WelcomeHandler)
	testName := "Bede"

	expectedStatusString := fmt.Sprintf("Welcome to private town, %s", testName)
	expectedStatus := test_app.Status{Status: expectedStatusString}

	// Excersising COT
	testStatus := testHandler.GetStatus(testName)

	// Asserting
	assert.Equal(t, expectedStatus, testStatus)

	// (Optional) Cleanup
}
