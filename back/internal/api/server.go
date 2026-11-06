package api

import (
	"fmt"
	"food/internal/api/routes"
	"log"
	"log/slog"
	"net/http"
)

var (
	SERVER_PORT = "3000"
)

func Serve() {

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", SERVER_PORT),
		Handler: routes.ConnectRoutes(),
	}
	log.Println(" -> server running on port :", SERVER_PORT)
	err := srv.ListenAndServe()

	if err != nil {
		slog.Error("error running server")
	}
}
