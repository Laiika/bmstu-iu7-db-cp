package main

import (
	"db_cp_6_sem/internal/app"
	"db_cp_6_sem/internal/config"
	"db_cp_6_sem/pkg/logger"
	_ "db_cp_6_sem/swagger"
)

//	@title			DB course project API
//	@version		1.0
//	@description	This is db course project backend API.

//	@contact.name	API Support
//	@contact.email	evgeniazavojskih@gmail.com

// @host		localhost:8080
// @BasePath	/api
// @Schemes	http
func main() {
	log := logger.GetLogger()
	cfg := config.GetConfig(log)

	app.Run(cfg, log)
}
