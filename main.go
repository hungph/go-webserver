package main

import (
	"./apps"
	"fmt"
	"os"
)

func main() {
	curApp := apps.App{}
	curApp.Initialize("")

	port := os.Getenv("port")

	if port == "" {
		port = "8080"
	}

	fmt.Print(port)

	curApp.Run(":" + port)
}
