package main

import (
	"fmt"
	"net/http"
	"zeronet-go/routers"
)

func main() {
	routers.InitRoutes()
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
