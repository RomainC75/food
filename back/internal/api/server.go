package api

import (
	"fmt"
	"food/config"
	"food/internal/api/routes"
	"log"
	"log/slog"
	"net/http"
)

func Serve() {
	port := config.Get().Server.Port
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes.ConnectRoutes(),
	}
	log.Println(" -> server running on port :", port)
	err := srv.ListenAndServe()

	if err != nil {
		slog.Error("error running server")
	}
}
