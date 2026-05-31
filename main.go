package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hx71/api-started-gin-golang/engine"
)

// @title Swagger for [Backend API Services]
// @version 1.0
// @description This is a document for API use in [Backend API Services]
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:1234
// @BasePath /
// @schemes http
func main() {
	r := engine.SetupRouter()

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}
}
