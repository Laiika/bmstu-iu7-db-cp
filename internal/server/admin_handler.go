package server

import (
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getAllUsers godoc
//
//	@Summary		Show all system users
//	@Description	return all system users
//	@Tags			admin
//	@Param			token  query	 string	 true	"User authentication token"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/users [get]
func (s *Server) getAllUsers(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	users, err := s.service.GetAllUsers(client, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"users": users})
}

// getUser godoc
//
//	@Summary		Show user with the specified id
//	@Description	return user with the specified id
//	@Tags			admin
//	@Param			token  query	 string	 true	"User authentication token"
//	@Param          id     path      int     true   "User id"
//	@Produce		json
//	@Success		200    {object}  map[string]interface{}
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		404	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/user/{id} [get]
func (s *Server) getUser(ctx *gin.Context) {
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

	user, err := s.service.GetUserById(client, ctx, id)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"user": user})
}

// createUser godoc
//
//	@Summary		Add new user to the system
//	@Description	add new user to the system
//	@Tags			admin
//	@Param			token	query	  string			 true	"User authentication token"
//	@Param			user    body	  entity.CreateUser	 true	"Information about stored user"
//	@Accept			json
//	@Produce		json
//	@Success		201         {object}  map[string]interface{}
//	@Failure		400	        {object}  map[string]interface{}
//	@Failure		500	        {object}  map[string]interface{}
//	@Router			/user [post]
func (s *Server) createUser(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var user entity.CreateUser
	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := s.service.CreateUser(client, ctx, &user)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

// updateUserRole godoc
//
//	@Summary		Update user role by id
//	@Description	update user role by id
//	@Tags			admin
//	@Param			token	query	    string	 true	"User authentication token"
//	@Param			role    body	    string	 true	"New user role"
//	@Param			id      path	    int 	 true	"User id"
//	@Accept			json
//	@Success		200
//	@Failure		400	    {object}    map[string]interface{}
//	@Failure		404	    {object}    map[string]interface{}
//	@Failure		500	    {object}    map[string]interface{}
//	@Router			/user/{id} [patch]
func (s *Server) updateUserRole(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := s.authService.GetClient(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var role string
	err = ctx.ShouldBindJSON(&role)
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.log.Error(err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = s.service.UpdateUserRole(client, ctx, id, role)
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

// deleteUser godoc
//
//	@Summary		Delete user by id
//	@Description	delete user by id
//	@Tags			admin
//	@Param			token  query	 string	 true   "User authentication token"
//	@Param          id     path      int     true   "User id"
//	@Success		200
//	@Failure		400	   {object}	 map[string]interface{}
//	@Failure		404	   {object}	 map[string]interface{}
//	@Failure		500	   {object}	 map[string]interface{}
//	@Router			/user/{id} [delete]
func (s *Server) deleteUser(ctx *gin.Context) {
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

	err = s.service.DeleteUser(client, ctx, id)
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
