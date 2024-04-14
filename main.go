package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"weather-app/weather"
)

func main() {
	// Ignore env load errors
	_ = godotenv.Load()

	r := gin.Default()

	weather.Route(r)

	err := r.Run() // 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
