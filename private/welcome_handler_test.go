package private

import (
	"github.com/kanerogers/test_app/test_app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWelcomeHandler(t *testing.T) {
	testHandler := new(WelcomeHandler)

	expectedStatus := test_app.Status{Status: "Welcome to private town"}

	assert.Equal(t, expectedStatus, testHandler.GetStatus())
}
