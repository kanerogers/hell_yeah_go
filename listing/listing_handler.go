package listing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Listing struct {
	ID   string
	Name string
}

type Handler struct {
}

func (*Handler) CreateHandler(c *gin.Context) {
	fmt.Println("Creating")
	c.JSON(http.StatusOK, gin.H{"status": "create"})
}

func (*Handler) TakeHandler(c *gin.Context) {
	fmt.Println("Taking")
	c.JSON(http.StatusOK, gin.H{"status": "take"})
}
