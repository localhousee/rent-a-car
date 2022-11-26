package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getAllCars)
	router.POST("/cars", storeCar)
	router.GET("/cars/:id", getSpecificCar)

	router.Run("localhost:8080")
}
