package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type circuit struct {
	ID       string `json:"id"`
	Provider string `json:"title"`
	Type     string `json:"author"`
	Speed    int    `json:"quanity"`
}

var circuits = []circuit{
	{"C_001", "Verizon", "optical", 100},
	{"C_002", "Spectrum", "optical", 50},
	{"C_003", "Comcast", "ethernet", 10},
}

func getCircuits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, circuits)
}

func main() {
	router := gin.Default()
	router.GET("/circuits", getCircuits)
	router.Run("localhost:8000")

}
