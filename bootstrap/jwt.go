package bootstrap

type JWTEnv struct {
	AccessTokenSecret string `env:"ACCESS_SECRET"`
	AccessTokenExpiry int64  `env:"ACCESS_EXPIRY"` // in seconds

	RefreshTokenSecret string `env:"REFRESH_SECRET"`
	RefreshTokenExpiry int64  `env:"REFRESH_EXPIRY"` // in seconds
}
