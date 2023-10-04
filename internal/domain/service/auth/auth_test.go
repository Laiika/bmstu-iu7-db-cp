package auth

import (
	"context"
	"db_cp_6_sem/internal/domain/entity"
	"db_cp_6_sem/internal/domain/service"
	"db_cp_6_sem/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestAuthService_Login(t *testing.T) {
	log := logger.GetLogger()
	var user, empl, admin service.Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	authService := NewAuthService(user, empl, admin, log)
	repo := service.NewMockIRepo(ctl)
	service := service.NewService(repo, log)

	name := "username"
	password := "11223"
	data := &entity.Auth{
		Username: name,
		Password: password,
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	assert.NoError(t, err)
	userCur := &entity.User{
		Id:       1,
		Name:     name,
		Password: string(bytes),
		Role:     "user",
	}

	repo.EXPECT().GetUserByName(admin, gomock.Any(), data.Username).Return(userCur, nil)
	_, err = authService.Login(context.Background(), service, data)
	assert.NoError(t, err)
}

func TestAuthService_Logout(t *testing.T) {
	log := logger.GetLogger()
	var user, empl, admin service.Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	authService := NewAuthService(user, empl, admin, log)
	repo := service.NewMockIRepo(ctl)
	service := service.NewService(repo, log)

	name := "username"
	password := "11223"
	data := &entity.Auth{
		Username: name,
		Password: password,
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	assert.NoError(t, err)
	userCur := &entity.User{
		Id:       1,
		Name:     name,
		Password: string(bytes),
		Role:     "user",
	}

	repo.EXPECT().GetUserByName(admin, gomock.Any(), data.Username).Return(userCur, nil)
	token, err := authService.Login(context.Background(), service, data)

	err = authService.Logout(token)
	assert.NoError(t, err)
}

func TestAuthService_GetSession(t *testing.T) {
	log := logger.GetLogger()
	var user, empl, admin service.Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	authService := NewAuthService(user, empl, admin, log)
	repo := service.NewMockIRepo(ctl)
	service := service.NewService(repo, log)

	name := "username"
	password := "11223"
	data := &entity.Auth{
		Username: name,
		Password: password,
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	assert.NoError(t, err)
	userCur := &entity.User{
		Id:       1,
		Name:     name,
		Password: string(bytes),
		Role:     "user",
	}

	repo.EXPECT().GetUserByName(admin, gomock.Any(), data.Username).Return(userCur, nil)
	token, _ := authService.Login(context.Background(), service, data)

	ok := authService.GetSession(token)
	assert.Equal(t, true, ok)
}

func TestAuthService_GetClient(t *testing.T) {
	log := logger.GetLogger()
	var user, empl, admin service.Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	authService := NewAuthService(user, empl, admin, log)
	repo := service.NewMockIRepo(ctl)
	service := service.NewService(repo, log)

	name := "username"
	password := "11223"
	data := &entity.Auth{
		Username: name,
		Password: password,
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	assert.NoError(t, err)
	userCur := &entity.User{
		Id:       1,
		Name:     name,
		Password: string(bytes),
		Role:     "user",
	}

	repo.EXPECT().GetUserByName(admin, gomock.Any(), data.Username).Return(userCur, nil)
	token, _ := authService.Login(context.Background(), service, data)

	client, err := authService.GetClient(token)
	assert.Equal(t, user, client)
	assert.NoError(t, err)
}
