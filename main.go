package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"weather-app/weather"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	weather.Route(r)

	err = r.Run() // 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
