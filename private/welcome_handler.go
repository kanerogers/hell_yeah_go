package private

import (
	"github.com/kanerogers/test_app/common"
)

type WelcomeHandler struct {
}

func (*WelcomeHandler) GetStatus() test_app.Status {
	return test_app.Status{Status: "Welcome to private town"}
}
