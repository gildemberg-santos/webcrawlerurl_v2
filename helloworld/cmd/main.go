package main

import (
	"log"

	_ "example.com/gcf"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.StartHostPort: %v\n", err)
	}
}
