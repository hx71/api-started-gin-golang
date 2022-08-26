package main

import (
	"os"

	"github.com/hasrulrhul/service-repository-pattern-gin-golang/engine"
)

func main() {

	r := engine.SetupRouter()

	r.Run(":" + os.Getenv("APP_PORT"))
}
