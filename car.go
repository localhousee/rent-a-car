package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type car struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Price int    `json:"price"`
}

var cars = []car{
	{ID: "1", Name: "Toyota", Type: "Corolla", Price: 1000},
	{ID: "2", Name: "Daihatsu", Type: "Xenia", Price: 500},
	{ID: "3", Name: "Honda", Type: "Jazz", Price: 200},
}

func getAllCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func storeCar(c *gin.Context) {
	var newCar car
	if err := c.BindJSON(&newCar); err != nil {

		return
	}
	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func getSpecificCar(c *gin.Context) {
	id := c.Param("id")

	for _, value := range cars {
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car not found"})
}
