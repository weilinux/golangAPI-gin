package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/golangAPI-gin/database"
	"github.com/weilinux/golangAPI-gin/middlewares"
	"github.com/weilinux/golangAPI-gin/src"
	"io"
	"os"
)

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


//gin 合适用来做restful api , 不适合mvc
func main() {
	setupLogging()
	router := gin.Default()
	router.Use(gin.Recovery(), middlewares.Logger())

	//http://localhost:8000/v1/users/
	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	//Connect to mysql server
	go func() {
		database.Connect()
	}()
/*
	//map path to handler function
	router.GET("/ping", func(c *gin.Context) {
		//gin.H gin.Header的意思
		c.JSON(200, gin.H{
			"message": "ping ok",
			"message2": "success",
		})

		
	})
	router.POST("/ping/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})

	})
	*/

	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
