package config

type MainConfig struct {
	PORT      string
	PORT_GRPC string

	JWT_SECRET_KEY string

	DB_PostgresqlHost     string
	DB_PostgresqlPort     string
	DB_PostgresqlUser     string
	DB_PostgresqlPassword string
	DB_PostgresqlDbname   string
	DB_PostgresqlSslmode  bool
	DB_PgDriver           string
	DB_maxOpenConns       int
	DB_connMaxLifetime    int
	DB_maxIdleConns       int
	DB_connMaxIdleTime    int
}
