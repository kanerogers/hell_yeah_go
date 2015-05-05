package private

import (
	"fmt"
	"github.com/kanerogers/test_app/common"
)

type WelcomeHandler struct {
}

func (*WelcomeHandler) GetStatus(name string) test_app.Status {
	greeting := fmt.Sprintf("Welcome to private town, %s", name)
	return test_app.Status{Status: greeting}
}
