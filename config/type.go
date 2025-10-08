package config

type MainConfig struct {
	PORT        string `mapstructure:"PORT"`
	PORT_GRPC   string `mapstructure:"PORT_GRPC"`
	SERVER_PORT string `mapstructure:"SERVER_PORT"`

	JWT_SECRET_KEY string `mapstructure:"JWT_SECRET_KEY"`

	DatabaseURL string `mapstructure:"DATABASE_URL"`
}
