package bootstrap

import "fmt"

type RedisEnv struct {
	Host     string `env:"HOST"`
	Port     uint   `env:"PORT"`
	Password string `env:"PASSWORD" envDefault:""`
}

func (env *RedisEnv) DSN() string {
	return fmt.Sprintf("%s:%d", env.Host, env.Port)
}
