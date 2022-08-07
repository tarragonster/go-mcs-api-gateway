package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tarragonster/go-mcs-api-gateway/pkg/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	r.Run(c.Port)
}
