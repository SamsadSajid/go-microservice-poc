package main

import (
	"github.com/gin-gonic/gin"
	"os"
)


func isDebugMode() bool {
	return gin.Mode() == os.Getenv("dev_env")
}

func isTestMode() bool {
	return gin.Mode() == os.Getenv("stage_env")
}

func isProdMode() bool {
	return gin.Mode() == os.Getenv("prod_env")
}

func isError(err error) bool {
	return err != nil
}
