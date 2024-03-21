package main

import (
	"net/http"

	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

type circuit struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Type     string `json:"type"`
	Speed    int    `json:"speed"`
}

var circuits = []circuit{
	{"C_001", "Verizon", "optical", 100},
	{"C_002", "Spectrum", "optical", 50},
	{"C_003", "Comcast", "ethernet", 10},
}

func getCircuits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, circuits)
}

func addCircuit(c *gin.Context) {
	var newCircuit circuit

	if err := c.BindJSON(&newCircuit); err != nil {
		return
	}

	circuits = append(circuits, newCircuit)
	c.IndentedJSON(http.StatusCreated, newCircuit)
}

func getCircuitByID(id string) (*circuit, error) {
	for i, c := range circuits {
		if c.ID == id {
			return &circuits[i], nil
		}
	}
	return nil, errors.New("circuit not found")
}

func circuitByID(c *gin.Context) {
	id := c.Param("id")
	circuit, err := getCircuitByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Circuit not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, circuit)

}

func updateCircuit(c *gin.Context) {
	id := c.Param("id")
	var updatedCircuit circuit
	if err := c.BindJSON(&updatedCircuit); err != nil {
		return
	}
	log.Print(updatedCircuit)

	for i, circuit := range circuits {
		if circuit.ID == id {
			if updatedCircuit.ID != "" {
				circuits[i].ID = updatedCircuit.ID
			}
			if updatedCircuit.Provider != "" {
				circuits[i].Provider = updatedCircuit.Provider
			}
			if updatedCircuit.Type != "" {
				circuits[i].Type = updatedCircuit.Type
			}
			if updatedCircuit.Speed != 0 {
				circuits[i].Speed = updatedCircuit.Speed
			}
			c.JSON(http.StatusOK, circuits[i])
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Circuit not found"})

}

func main() {
	router := gin.Default()
	router.GET("/circuits", getCircuits)
	router.GET("/circuits/:id", circuitByID)
	router.POST("/circuits", addCircuit)
	router.PATCH("/circuits/:id", updateCircuit)
	router.Run("localhost:8000")

}
