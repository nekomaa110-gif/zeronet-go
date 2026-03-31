package routers

import (
	"net/http"
	"zeronet-go/handlers"
)

func initRoutes() {
	http.HandleFunc("/", handlers.Home)
}
