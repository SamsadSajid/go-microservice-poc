package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init(){
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	gin.SetMode(os.Getenv("stage_env"))
}

func main(){
	router := Controller()

	router.Run(":" + os.Getenv("PORT"))
}
