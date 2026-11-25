package config

import (
	"food/utils/displayer"
	"strings"

	"github.com/spf13/viper"
)

func Set() {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	configurations = Config{
		App: App{
			Name: v.GetString("APP_NAME"),
		},
		Server: Server{
			Port: v.GetString("BACK_PORT"),
		},
		Jwt: Jwt{
			RefreshSecret:        v.GetString("JWT_REFRESH_SECRET"),
			AccessSecret:         v.GetString("JWT_ACCESS_SECRET"),
			RefreshExpiresInDays: v.GetInt("JWT_REFRESH_EXPIRES_IN_DAYS"),
			AccessExpiresInMin:   v.GetInt("JWT_ACCESS_EXPIRES_IN_MIN"),
		},
	}

	displayer.PrettyDisplay("config", configurations)
}
