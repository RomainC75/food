package bootstrap

import (
	"food/config"
	"food/internal/api"
)

func Run() {
	config.Set()

	api.Serve()
}
