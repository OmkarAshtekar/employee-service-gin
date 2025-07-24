package main

import (
	"employee-service-gin/router"
	"fmt"
)

func main() {
	r := router.SetupRouter()
	fmt.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
