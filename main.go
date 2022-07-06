package main

import (
	"fmt"
	"kunimitsuh/app-runner-test/connection"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type JsonResp struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	router.GET("/api", func(c *gin.Context) {
		var j JsonResp
		j.A = 10
		j.B = 20
		c.JSON(http.StatusOK, j)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/health/check/network", func(c *gin.Context) {
		_, err := http.Get("http://clients3.google.com/generate_204")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "internet connection error: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	router.GET("/health/check/deep", func(c *gin.Context) {
		db, err := connection.DbInit()
		if err != nil {
			c.String(http.StatusInternalServerError, "DB connection error")
			return
		}
		c.String(http.StatusOK, "ok")
	})

	return router
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnvFile()
	router := setupRouter()
	router.Run(":8080")
}
