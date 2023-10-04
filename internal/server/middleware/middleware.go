package middleware

import (
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/service/auth"
	"github.com/gin-gonic/gin"
	pkgErrors "github.com/pkg/errors"
	"net/http"
)

func SessionCheck(service *auth.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")

		if !service.GetSession(token) {
			ctx.JSON(http.StatusBadRequest, pkgErrors.WithMessage(apperror.ErrSessionNotExists, token))
			return
		}

		ctx.Next()
	}
}
