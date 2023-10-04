package server

import (
	"db_cp_6_sem/internal/apperror"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

// getAllAnalyzers godoc
//
//	@Summary		Show all analyzers
//	@Description	return all analyzers
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/analyzers [get]
func (s *Server) getAllAnalyzers(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	analyzers, err := s.service.GetAllAnalyzers(client, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"analyzers": analyzers})
}

// getAnalyzersByType godoc
//
//	@Summary		Show analyzers of specified type
//	@Description	return all analyzers of specified type
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Param          name   path      string  true   "Name of analyzer type"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/analyzer_type/{name}/analyzers [get]
func (s *Server) getAnalyzersByType(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	name, _ := strings.CutPrefix(ctx.Param("name"), "/")
	analyzers, err := s.service.GetTypeAnalyzers(client, ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"analyzers": analyzers})
}

// getAnalyzer godoc
//
//	@Summary		Show analyzer with the specified id
//	@Description	return analyzer with the specified id
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Param          id     path      string  true   "Analyzer id"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		404	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/analyzer/{id} [get]
func (s *Server) getAnalyzer(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, _ := strings.CutPrefix(ctx.Param("id"), "/")
	analyzer, err := s.service.GetAnalyzerById(client, ctx, id)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"analyzer": analyzer})
}

// getAllTypes godoc
//
//	@Summary		Show all analyzer types
//	@Description	return all analyzer types
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/analyzer_types [get]
func (s *Server) getAllTypes(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	types, err := s.service.GetAllTypes(client, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"types": types})
}

// getAllSensors godoc
//
//	@Summary		Show all sensors
//	@Description	return all sensors
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/sensors [get]
func (s *Server) getAllSensors(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	sensors, err := s.service.GetAllSensors(client, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"sensors": sensors})
}

// getSensorsByAnalyzerId godoc
//
//	@Summary		Show sensors of specified analyzer
//	@Description	return all sensors of specified analyzer
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Param          id     path      string  true   "Analyzer id"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/analyzer/{id}/sensors [get]
func (s *Server) getSensorsByAnalyzerId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, _ := strings.CutPrefix(ctx.Param("id"), "/")
	sensors, err := s.service.GetAnalyzerSensors(client, ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"sensors": sensors})
}

// getEventsBySensorId godoc
//
//	@Summary		Show events of specified sensor
//	@Description	return all events of specified sensor
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Param          id     path      string  true   "Sensor id"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/sensor/{id}/events [get]
func (s *Server) getEventsBySensorId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, _ := strings.CutPrefix(ctx.Param("id"), "/")
	events, err := s.service.GetSensorEvents(client, ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"events": events})
}

// getEventsBySignalTime godoc
//
//	@Summary		Show events that occurred during the specified period of time
//	@Description	return all events that occurred during the specified period of time
//	@Tags			user
//	@Param			token   query	  string	 true	"User authentication token"
//	@Param          start   path      string  true   "Left period border"
//	@Param          finish  path      string  true   "Right period border"
//	@Produce		json
//	@Success		200     {object}  map[string]interface{}
//	@Failure		400	    {object}  map[string]interface{}
//	@Failure		500	    {object}  map[string]interface{}
//	@Router			/events/{start}/{finish} [get]
func (s *Server) getEventsBySignalTime(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	st, _ := strings.CutPrefix(ctx.Param("start"), "/")
	st, _ = strings.CutSuffix(st, "/")
	start, err := time.Parse("2006-01-02 03:04:05", st)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	fin, _ := strings.CutPrefix(ctx.Param("finish"), "/")
	finish, err := time.Parse("2006-01-02 03:04:05", fin)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	events, err := s.service.GetEventsBySignalTime(client, ctx, start, finish)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"events": events})
}

// getAllGases godoc
//
//	@Summary		Show all gases
//	@Description	return all gases
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/gases [get]
func (s *Server) getAllGases(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	gases, err := s.service.GetAllGases(client, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"gases": gases})
}

// getGasesByType godoc
//
//	@Summary		Show gases of specified type
//	@Description	return all gases of specified type
//	@Tags			user
//	@Param			token  query	 string	 true	"User authentication token"
//	@Param          name   path      string  true   "Name of analyzer type"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/analyzer_type/{name}/gases [get]
func (s *Server) getGasesByType(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	name, _ := strings.CutPrefix(ctx.Param("name"), "/")
	gases, err := s.service.GetTypeGases(client, ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"gases": gases})
}
