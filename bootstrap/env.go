package bootstrap

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type Env struct {
	DB     DBEnv    `envPrefix:"DB_"`
	Redis  RedisEnv `envPrefix:"REDIS_"`
	Server Server   `envPrefix:"SERVER_"`
	JWT    JWTEnv   `envPrefix:"JWT_"`
}

func NewEnv() *Env {
	var e Env
	if err := env.ParseWithOptions(&e, env.Options{
		RequiredIfNoDef: true,
		Prefix:          "BOOKSTORE_",
	}); err != nil {
		log.Fatal(err)
	}
	return &e
}
