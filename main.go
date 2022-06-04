package main

import (
	"kunimitsuh/app-runner-test/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResp struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
}


func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	router.GET("/api", func(c *gin.Context){
		var j JsonResp
		j.A = 10
		j.B = 20
		c.JSON(http.StatusOK, j)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/health/check/deep", func(c *gin.Context) {
		db, err := connection.DbInit()
		if err != nil {
			c.String(http.StatusInternalServerError, "DB connection error")
			return
		}
		c.String(http.StatusOK, db.Name())
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
