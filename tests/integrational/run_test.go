package integrational

import (
	"context"
	"db_cp_6_sem/internal/config"
	"db_cp_6_sem/internal/db/postgres"
	"db_cp_6_sem/internal/domain/service"
	"db_cp_6_sem/pkg/client/postgresql"
	"db_cp_6_sem/pkg/logger"
	"os"
	"testing"
)

var (
	client postgresql.Client
	srvc   *service.Service
)

func setup() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	var err error
	client, err = postgresql.NewClient(context.Background(), 3, &cfg.Admin)
	if err != nil {
		log.Fatal(err)
	}

	repo := postgres.NewPostgresRepo()
	srvc = service.NewService(repo, log)
}

func shutdown() {
	client.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
