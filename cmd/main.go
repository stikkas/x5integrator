package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stikkas/integrator/internal/handler"
	"log"
)

func main() {
	engine := gin.Default()
	handler.Create(engine)
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
