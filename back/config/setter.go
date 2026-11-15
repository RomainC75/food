package config

import (
	"log"
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
			Secret:       v.GetString("JWT_SECRET"),
			ExpiresInMin: v.GetInt("JWT_EXPIRES_IN_MIN"),
		},
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}

}
