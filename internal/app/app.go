package app

import (
	"context"
	"db_cp_6_sem/internal/config"
	"db_cp_6_sem/internal/db/postgres"
	"db_cp_6_sem/internal/domain/service"
	"db_cp_6_sem/internal/domain/service/auth"
	"db_cp_6_sem/internal/server"
	"db_cp_6_sem/pkg/client/postgresql"
	"db_cp_6_sem/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config, log *logger.Logger) {
	user, err := postgresql.NewClient(context.Background(), 3, &cfg.User)
	if err != nil {
		log.Fatal(err)
	}
	defer user.Close()

	empl, err := postgresql.NewClient(context.Background(), 3, &cfg.Empl)
	if err != nil {
		log.Fatal(err)
	}
	defer empl.Close()

	admin, err := postgresql.NewClient(context.Background(), 3, &cfg.Admin)
	if err != nil {
		log.Fatal(err)
	}
	defer admin.Close()
	log.Info("connected to db")

	repo := postgres.NewPostgresRepo()
	service := service.NewService(repo, log)

	authService := auth.NewAuthService(user, empl, admin, log)

	srv := server.NewServer(&cfg.Server, authService, service, log)
	go func() {
		err = srv.Start()
		if err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("listen: %s\n", err))
		}
	}()
	log.Info("server started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err = srv.Stop()
	if err != nil {
		log.Errorf("server shutdown: %v", err)
	}
	log.Info("server exited")
}
