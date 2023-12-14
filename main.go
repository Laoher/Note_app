package main

import (
	"gin_notes/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/**/*")

	models.ConnectDB()
	models.DBMigrate()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes application",
			"content": "HelloWorld!",
		})
	})

	log.Println("Server Started!")
	r.Run() // Default Port 8080
}
