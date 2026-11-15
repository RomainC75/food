package api

import (
	"fmt"
	"food/config"
	"food/internal/api/routes"
	"log/slog"
	"net/http"
	"time"
)

func Serve() {
	port := config.Get().Server.Port
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes.ConnectRoutes(),
	}
	slog.Info(" -> server running on port : ", port, time.Now())
	err := srv.ListenAndServe()

	if err != nil {
		slog.Error("error running server")
	}
}
