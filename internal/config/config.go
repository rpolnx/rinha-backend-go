package configs

type EnvConfig struct {
	Port       int    `env:"PORT" envDefault:"8080"`
	DbHost     string `env:"DB_HOST"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbDbname   string `env:"DB_DBNAME"`
	DbPort     int    `env:"DB_PORT"`
	DbSslmode  string `env:"DB_SSLMODE" envDefault:"disable"`
	DbTimezone string `env:"DB_TIMEZONE" envDefault:"DB_TIMEZONE"`
}
