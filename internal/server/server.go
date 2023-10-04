package server

import (
	"db_cp_6_sem/internal/config"
	"db_cp_6_sem/internal/domain/entity"
	"db_cp_6_sem/internal/domain/service"
	"db_cp_6_sem/internal/domain/service/auth"
	"db_cp_6_sem/internal/server/middleware"
	"db_cp_6_sem/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Server struct {
	srv         *http.Server
	authService *auth.AuthService
	service     *service.Service
	log         *logger.Logger
}

// login godoc
//
//	@Summary		Log in to the server
//	@Description	log in to the server
//	@Tags			common
//	@Param			data  body	    entity.Auth	 true	"Authentication request"
//	@Accept			json
//	@Produce		json
//	@Success		200	  {object}  map[string]interface{}
//	@Failure		400	  {object}  map[string]interface{}
//	@Router			/login [post]
func (s *Server) login(ctx *gin.Context) {
	var data entity.Auth
	if err := ctx.ShouldBindJSON(&data); err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	token, err := s.authService.Login(ctx, s.service, &data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"Token": token})
}

// logout godoc
//
//	@Summary		Log out from the server
//	@Description	log out from the server
//	@Tags			common
//	@Param			token  query	 string	 true	"User authentication token"
//	@Success		200
//	@Failure		400	   {object}	 map[string]interface{}
//	@Router			/logout [post]
func (s *Server) logout(ctx *gin.Context) {
	token := ctx.Query("token")

	err := s.authService.Logout(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) registerRoutes(gr *gin.RouterGroup) {
	gr.GET("/analyzers", s.getAllAnalyzers)
	gr.GET("/analyzer_type/:name/analyzers", s.getAnalyzersByType)
	gr.GET("/analyzer/:id", s.getAnalyzer)
	gr.POST("/analyzer", s.createAnalyzer)
	gr.DELETE("/analyzer/:id", s.deleteAnalyzer)

	gr.GET("/analyzer_types", s.getAllTypes)
	gr.POST("/analyzer_type", s.createType) // указать все газы
	gr.DELETE("/analyzer_type/:name", s.deleteType)

	gr.GET("/sensors", s.getAllSensors)
	gr.GET("/analyzer/:id/sensors", s.getSensorsByAnalyzerId)
	gr.POST("/sensor", s.createSensor)
	gr.PATCH("/sensor/:id", s.updateSensorAnalyzer)
	gr.DELETE("/sensor/:id", s.deleteSensor)

	gr.GET("/sensor/:id/events", s.getEventsBySensorId)
	gr.GET("/events/:start/:finish", s.getEventsBySignalTime)
	gr.POST("/event", s.createEvent)
	gr.DELETE("/event/:id", s.deleteEvent)

	gr.GET("/gases", s.getAllGases)
	gr.GET("/analyzer_type/:name/gases", s.getGasesByType)
	gr.POST("/gas", s.createGas)

	gr.GET("/users", s.getAllUsers)
	gr.GET("/user/:id", s.getUser)
	gr.POST("/user", s.createUser)
	gr.PATCH("/user/:id", s.updateUserRole)
	gr.DELETE("/user/:id", s.deleteUser)
}

func NewServer(cfg *config.ServerConfig, authService *auth.AuthService, service *service.Service, log *logger.Logger) *Server {
	gin.DisableConsoleColor()

	router := gin.Default()
	router.Use(gin.LoggerWithWriter(log.Writer()))
	router.Use(gin.RecoveryWithWriter(log.Writer()))

	//url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", cfg.Host, cfg.Port))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s := Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Handler: router,
		},
		authService: authService,
		service:     service,
		log:         log,
	}

	mainGroup := router.Group("/api")
	mainGroup.POST("/login", s.login)
	mainGroup.POST("/logout", s.logout)

	withAuth := mainGroup.Group("", middleware.SessionCheck(s.authService))
	s.registerRoutes(withAuth)

	return &s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
