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

func (*Listing) CreateHandler(c *gin.Context) {
	fmt.Println("Creating")
	c.JSON(http.StatusOK, gin.H{"status": "create"})
}

func (*Listing) TakeHandler(c *gin.Context) {
	fmt.Println("Taking")
	c.JSON(http.StatusOK, gin.H{"status": "take"})
}
