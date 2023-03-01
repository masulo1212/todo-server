package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Config struct {
	router *gin.Engine
}

func main() {
	r := gin.Default()

	app := Config{
		router: r,
	}

	app.Routes()

	err := r.Run(":80")
	if err != nil {
		log.Println("cannot connect to :80")
	}

}
