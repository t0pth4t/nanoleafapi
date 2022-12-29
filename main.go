package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var router *gin.Engine
var client *http.Client
var IP string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	IP = os.Getenv("IP")
	client = &http.Client{}
	router = gin.Default()
	router.Use(corsMiddleware())
	initializeRoutes()
	err = router.Run()
	if err != nil {
		return
	}
}
