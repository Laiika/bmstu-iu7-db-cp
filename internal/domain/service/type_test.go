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

func TestService_GetAllTypes(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	expectedTypes := entity.Types{
		&entity.Type{
			Name:       "VENTIS MX4",
			MaxSensors: 4,
		},
		&entity.Type{
			Name:       "VENTIS MX6",
			MaxSensors: 6,
		},
	}

	repo.EXPECT().GetAllTypes(client, gomock.Any()).Return(expectedTypes, nil)
	types, err := service.GetAllTypes(client, context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedTypes, types)
}

func TestService_CreateType(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	name := "VENTIS MX4"
	gases := []string{"кислород", "метан"}
	anType := &entity.CreateType{
		Name:       name,
		MaxSensors: 3,
		Gases:      gases,
	}
	tp := &entity.Type{
		Name:       name,
		MaxSensors: 3,
	}

	repo.EXPECT().GetTypeByName(client, gomock.Any(), name).Return(nil, apperror.ErrEntityNotFound)
	repo.EXPECT().CreateType(client, gomock.Any(), tp, gases).Return(nil)
	err := service.CreateType(client, context.Background(), anType)
	assert.NoError(t, err)
}

func TestService_DeleteType(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := "20081UC-061"
	repo.EXPECT().DeleteType(client, gomock.Any(), id).Return(nil)
	err := service.DeleteType(client, context.Background(), id)
	assert.NoError(t, err)
}
