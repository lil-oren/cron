package dependency

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		App        app
		PostgreDB  postgreDB
		RedisCache redisconfig
	}

	app struct {
		GracefulTimeout uint   `env:"GRACEFUL_TIMEOUT"`
	}

	postgreDB struct {
		DBHost string `env:"DB_HOST"`
		DBPort string `env:"DB_PORT"`
		DBName string `env:"DB_NAME"`
		DBUser string `env:"DB_USER"`
		DBPass string `env:"DB_PASS"`
		DBTz   string `env:"DB_TZ" env-default:"Asia/Jakarta"`
	}

	redisconfig struct {
		HOST     string `env:"REDIS_HOST"`
		PORT     string `env:"REDIS_PORT"`
		Password string `env:"REDIS_PASSWORD"`
	}
)

func NewConfig(logger Logger) (*Config, error) {
	config := new(Config)

	err := cleanenv.ReadEnv(config)
	if err != nil {
		logger.Fatalf("Failed to load config")
		return nil, err
	}

	logger.Infof("Successfully load config", nil)

	return config, err
}
