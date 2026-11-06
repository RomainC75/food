package bootstrap

import (
	"food/config"
	"food/internal/api"
	api_validator "food/internal/api/validator"
)

func Run() {
	config.Set()
	api_validator.Set()

	api.Serve()
}
