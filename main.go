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

	// dbEvent := os.Getenv("DBEVENT")
	// if dbEvent == "rollback" {
	// 	migration.RunRollback()
	// } else if dbEvent == "migration" {
	// 	migration.RunMigrations()
	// } else if dbEvent == "seeder" {
	// 	migration.RunMigrations()
	// 	seeder.RunSeeder()
	// }

	r := engine.SetupRouter()
	r.Run(":" + os.Getenv("APP_PORT"))

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}
}
