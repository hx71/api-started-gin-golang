package main

import (
	"os"

	"github.com/hasrulrhul/service-repository-pattern-gin-golang/engine"
)

// @title Swagger for [Backend API Services]
// @version 1.0
// @description This is a document for API use in [Backend API Services]
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:1234
// @BasePath /api/v1
// @schemes http
func main() {

	r := engine.SetupRouter()

	r.Run(":" + os.Getenv("APP_PORT"))
}
