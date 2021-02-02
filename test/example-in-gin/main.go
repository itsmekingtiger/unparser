package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	r.GET("/")
	r.GET("/user")
	r.GET("/user/")

	r.Run(":5050")
}
