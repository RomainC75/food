package config

type Config struct {
	App    App
	Server Server
	Jwt    Jwt
}

type App struct {
	Name string
}

type Server struct {
	Port string
}

type Jwt struct {
	Secret       string
	ExpiresInMin int
}
