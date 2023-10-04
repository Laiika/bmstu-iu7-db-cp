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

func TestService_GetAllGases(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	expectedGases := entity.Gases{
		&entity.Gas{
			Name:    "метан",
			Formula: "H2O",
			Type:    "горючие газы",
		},
		&entity.Gas{
			Name:    "сероводород",
			Formula: "H2O",
			Type:    "токсичные газы",
		},
	}

	repo.EXPECT().GetAllGases(client, gomock.Any()).Return(expectedGases, nil)
	gases, err := service.GetAllGases(client, context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedGases, gases)
}

func TestService_GetTypeGases(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	anType := "VENTIS MX4"
	expectedGases := entity.Gases{
		&entity.Gas{
			Name:    "метан",
			Formula: "H2O",
			Type:    "горючие газы",
		},
		&entity.Gas{
			Name:    "2",
			Formula: "H2O",
			Type:    "токсичные газы",
		},
	}

	repo.EXPECT().GetTypeGases(client, gomock.Any(), anType).Return(expectedGases, nil)
	gases, err := service.GetTypeGases(client, context.Background(), anType)
	assert.NoError(t, err)
	assert.Equal(t, expectedGases, gases)
}

func TestService_CreateGas(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	name := "name"
	gas := &entity.Gas{
		Name:    name,
		Formula: "H2O",
		Type:    "горючие газы",
	}

	repo.EXPECT().GetGasByName(client, gomock.Any(), name).Return(nil, apperror.ErrEntityNotFound)
	repo.EXPECT().CreateGas(client, gomock.Any(), gas).Return(nil)
	err := service.CreateGas(client, context.Background(), gas)
	assert.NoError(t, err)
}
