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

	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
