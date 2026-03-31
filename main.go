package main

import (
	"fmt"
	"net/http"
	"zeronet-go/routers"
)

func main() {
	routers.InitRoutes()
	fmt.Println("Server Jalan di :8080")
	http.ListenAndServe(":8080", nil)
}
