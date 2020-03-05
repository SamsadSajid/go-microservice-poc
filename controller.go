package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func Controller() *gin.Engine{
	router := gin.Default()

	router.Use(gin.LoggerWithFormatter(Logger))

	if isTestMode() || isProdMode(){
		logFile, err := os.Create("out.log")

		if isError(err) {
			log.Fatal("Can not create output log")
		}
		gin.DefaultWriter = io.MultiWriter(logFile)
	}

	if isDebugMode(){
		gin.ForceConsoleColor()
	}

	router.Use(gin.Recovery())

	fmt.Println(gin.Mode())

	//router.GET("/ping", func(c *gin.Context) {
	//	log.SetOutput()
	//
	//	panic("A problem")
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "PONG",
	//	})
	//})


	//router.GET("/service/:userName", getUserInfo)

	router.POST("/service/api/rabbit", getData)

	return router
}


