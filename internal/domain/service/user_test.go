package service

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"db_cp_6_sem/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetUserById(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := 3
	expectedUser := &entity.User{
		Id:       id,
		Name:     "jfdjkfd",
		Password: "jfdjkdf",
		Role:     "admin",
	}

	repo.EXPECT().GetUserById(client, gomock.Any(), id).Return(expectedUser, nil)
	user, err := service.GetUserById(client, context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestService_GetUserByName(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	name := "kjdkd"
	expectedUser := &entity.User{
		Id:       1,
		Name:     name,
		Password: "jfdjkdf",
		Role:     "admin",
	}

	repo.EXPECT().GetUserByName(client, gomock.Any(), name).Return(expectedUser, nil)
	user, err := service.GetUserByName(client, context.Background(), name)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestService_GetAllUsers(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	expectedUsers := entity.Users{
		&entity.User{
			Id:       1,
			Name:     "jfdjkfd",
			Password: "jfdjkdf",
			Role:     "admin",
		},
		&entity.User{
			Id:       5,
			Name:     "jfdjkfd",
			Password: "jfdjkdf",
			Role:     "user",
		},
	}

	repo.EXPECT().GetAllUsers(client, gomock.Any()).Return(expectedUsers, nil)
	users, err := service.GetAllUsers(client, context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestService_CreateUser(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	name := "jfdjkfd"
	us := &entity.CreateUser{
		Name:     name,
		Password: "jfdjkdf",
		Role:     "user",
	}

	id := 4
	repo.EXPECT().GetUserByName(client, gomock.Any(), name).Return(nil, apperror.ErrEntityNotFound)
	repo.EXPECT().CreateUser(client, gomock.Any(), gomock.Any()).Return(id, nil)
	id2, err := service.CreateUser(client, context.Background(), us)
	assert.NoError(t, err)
	assert.Equal(t, id, id2)
}

func TestService_UpdateUserRole(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := 5
	role := "employee"
	repo.EXPECT().UpdateUserRole(client, gomock.Any(), id, role).Return(nil)
	err := service.UpdateUserRole(client, context.Background(), id, role)
	assert.NoError(t, err)
}

func TestService_DeleteUser(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := 3
	repo.EXPECT().DeleteUser(client, gomock.Any(), id).Return(nil)
	err := service.DeleteUser(client, context.Background(), id)
	assert.NoError(t, err)
}
