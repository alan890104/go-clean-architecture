package bootstrap

type Server struct {
	Port     uint   `env:"PORT"`
	TimeZone string `env:"TIMEZONE"`
}
