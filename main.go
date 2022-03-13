package main

import (
	"os"
	"srp-golang/engine"
)

func main() {

	r := engine.SetupRouter()

	r.Run(":" + os.Getenv("APP_PORT"))
}
