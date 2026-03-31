package routers

import (
	"net/http"
	"zeronet-go/handlers"
)

func InitRoutes() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/test", handlers.Test)
}
