package main

import (
	"github.com/lil-oren/cron/internal/dependency"
	"github.com/lil-oren/cron/internal/infra"
)

func main() {
	logger := dependency.NewLogger()

	config, err := dependency.NewConfig(logger)
	if err != nil {
		return
	}

	db, err := dependency.NewPGDB(*config, logger)
	if err != nil {
		return
	}

	rc := dependency.NewRedisClient(*config, logger)
	if rc == nil {
		return
	}

	infra.InitApp(db, rc, *config, logger)
}
