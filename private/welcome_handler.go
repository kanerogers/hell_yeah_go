package private

import (
	"github.com/kanerogers/test_app/test_app"
)

type WelcomeHandler struct {
}

func (*WelcomeHandler) GetStatus() test_app.Status {
	return test_app.Status{Status: "Welcome to private town"}
}
