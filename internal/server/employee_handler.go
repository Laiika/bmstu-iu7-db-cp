package server

import (
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

// createAnalyzer godoc
//
//	@Summary		Add new analyzer
//	@Description	add new analyzer
//	@Tags			employee
//	@Param			token		query	  string			 true	"User authentication token"
//	@Param			analyzer    body	  entity.Analyzer	 true	"Information about stored analyzer"
//	@Accept			json
//	@Success		201
//	@Failure		400	        {object}  map[string]interface{}
//	@Failure		500	        {object}  map[string]interface{}
//	@Router			/analyzer [post]
func (s *Server) createAnalyzer(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var analyzer entity.Analyzer
	err = ctx.ShouldBindJSON(&analyzer)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = s.service.CreateAnalyzer(client, ctx, &analyzer)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

// deleteAnalyzer godoc
//
//	@Summary		Delete analyzer by id
//	@Description	delete analyzer by id
//	@Tags			employee
//	@Param			token  query	 string	 true   "User authentication token"
//	@Param          id     path      string  true   "Analyzer id"
//	@Success		200
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		404	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/analyzer/{id} [delete]
func (s *Server) deleteAnalyzer(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, _ := strings.CutPrefix(ctx.Param("id"), "/")
	err = s.service.DeleteAnalyzer(client, ctx, id)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// createType godoc
//
//	@Summary		Add new analyzer type
//	@Description	add new analyzer type
//	@Tags			employee
//	@Param			token	  query	    string			     true	"User authentication token"
//	@Param			anType    body	    entity.CreateType	 true	"Information about stored analyzer type"
//	@Accept			json
//	@Success		201
//	@Failure		400	      {object}  map[string]interface{}
//	@Failure		500	      {object}  map[string]interface{}
//	@Router			/analyzer_type [post]
func (s *Server) createType(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var anType entity.CreateType
	err = ctx.ShouldBindJSON(&anType)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = s.service.CreateType(client, ctx, &anType)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

// deleteType godoc
//
//	@Summary		Delete analyzer type by name
//	@Description	delete analyzer type by name
//	@Tags			employee
//	@Param			token  query	 string	 true   "User authentication token"
//	@Param          name   path      string  true   "Analyzer type name"
//	@Success		200
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		404	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/analyzer_type/{name} [delete]
func (s *Server) deleteType(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	name, _ := strings.CutPrefix(ctx.Param("name"), "/")
	err = s.service.DeleteType(client, ctx, name)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// createSensor godoc
//
//	@Summary		Add new sensor
//	@Description	add new sensor
//	@Tags			employee
//	@Param			token	  query	    string			 true	"User authentication token"
//	@Param			sensor    body	    entity.Sensor	 true	"Information about stored sensor"
//	@Accept			json
//	@Success		201
//	@Failure		400	      {object}  map[string]interface{}
//	@Failure		500	      {object}  map[string]interface{}
//	@Router			/sensor [post]
func (s *Server) createSensor(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var sensor entity.Sensor
	err = ctx.ShouldBindJSON(&sensor)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = s.service.CreateSensor(client, ctx, &sensor)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

// updateSensorAnalyzer godoc
//
//	@Summary		Update sensor analyzer_id by id
//	@Description	update sensor analyzer_id by id
//	@Tags			employee
//	@Param			token	query	    string	 true	"User authentication token"
//	@Param			anId    body	    string	 true	"New analyzer id"
//	@Param			id      path	    string	 true	"Sensor id"
//	@Accept			json
//	@Success		200
//	@Failure		400	    {object}    map[string]interface{}
//	@Failure		404	    {object}    map[string]interface{}
//	@Failure		500	    {object}    map[string]interface{}
//	@Router			/sensor/{id} [patch]
func (s *Server) updateSensorAnalyzer(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var anId string
	err = ctx.ShouldBindJSON(&anId)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, _ := strings.CutPrefix(ctx.Param("id"), "/")
	err = s.service.UpdateSensorAnalyzer(client, ctx, id, anId)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// deleteSensor godoc
//
//	@Summary		Delete sensor by id
//	@Description	delete sensor by id
//	@Tags			employee
//	@Param			token  query	 string	 true   "User authentication token"
//	@Param          id     path      string  true   "Sensor id"
//	@Success		200
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		404	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/sensor/{id} [delete]
func (s *Server) deleteSensor(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, _ := strings.CutPrefix(ctx.Param("id"), "/")
	err = s.service.DeleteSensor(client, ctx, id)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// createEvent godoc
//
//	@Summary		Add new event
//	@Description	add new event
//	@Tags			employee
//	@Param			token	  query	    string			    true	"User authentication token"
//	@Param			event     body	    entity.CreateEvent	 true	"Information about stored event"
//	@Accept			json
//	@Produce		json
//	@Success		201       {object}  map[string]interface{}
//	@Failure		400	      {object}  map[string]interface{}
//	@Failure		500	      {object}  map[string]interface{}
//	@Router			/event [post]
func (s *Server) createEvent(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var event entity.CreateEvent
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := s.service.CreateEvent(client, ctx, &event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

// deleteEvent godoc
//
//	@Summary		Delete event by id
//	@Description	delete event by id
//	@Tags			employee
//	@Param			token  query	 string	 true   "User authentication token"
//	@Param          id     path      int     true   "Event id"
//	@Success		200
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		404	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/event/{id} [delete]
func (s *Server) deleteEvent(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = s.service.DeleteEvent(client, ctx, id)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// createGas godoc
//
//	@Summary		Add new gas
//	@Description	add new gas
//	@Tags			employee
//	@Param			token  query  string		 true	"User authentication token"
//	@Param			gas    body	  entity.Gas	 true	"Information about stored gas"
//	@Accept			json
//	@Success		201
//	@Failure		400	        {object}  map[string]interface{}
//	@Failure		500	        {object}  map[string]interface{}
//	@Router			/gas [post]
func (s *Server) createGas(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var gas entity.Gas
	err = ctx.ShouldBindJSON(&gas)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = s.service.CreateGas(client, ctx, &gas)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
