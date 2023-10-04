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

func TestService_GetAllSensors(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	expectedSensors := entity.Sensors{
		&entity.Sensor{
			Id:              "14032ER395",
			Type:            "электрохимический",
			AnalyzerId:      "20081UC-061",
			Gas:             "кислород",
			LowLimitAlarm:   "6 промилле",
			UpperLimitAlarm: "7 промилле",
		},
		&entity.Sensor{
			Id:              "20081UC-023",
			Type:            "электрохимический",
			AnalyzerId:      "20081UC-061",
			Gas:             "метан",
			LowLimitAlarm:   "10 % НКПВ",
			UpperLimitAlarm: "50 % НКПВ",
		},
	}

	repo.EXPECT().GetAllSensors(client, gomock.Any()).Return(expectedSensors, nil)
	sensors, err := service.GetAllSensors(client, context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedSensors, sensors)
}

func TestService_GetAnalyzerSensors(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	anId := "20081UC-061"
	expectedSensors := entity.Sensors{
		&entity.Sensor{
			Id:              "14032ER395",
			Type:            "электрохимический",
			AnalyzerId:      anId,
			Gas:             "кислород",
			LowLimitAlarm:   "6 промилле",
			UpperLimitAlarm: "7 промилле",
		},
		&entity.Sensor{
			Id:              "20081UC-023",
			Type:            "электрохимический",
			AnalyzerId:      anId,
			Gas:             "метан",
			LowLimitAlarm:   "10 % НКПВ",
			UpperLimitAlarm: "50 % НКПВ",
		},
	}

	repo.EXPECT().GetAnalyzerSensors(client, gomock.Any(), anId).Return(expectedSensors, nil)
	sensors, err := service.GetAnalyzerSensors(client, context.Background(), anId)
	assert.NoError(t, err)
	assert.Equal(t, expectedSensors, sensors)
}

func TestService_CreateSensor(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := "20081UC-023"
	sensor := &entity.Sensor{
		Id:              id,
		Type:            "электрохимический",
		AnalyzerId:      "20081UC-061",
		Gas:             "метан",
		LowLimitAlarm:   "10 % НКПВ",
		UpperLimitAlarm: "50 % НКПВ",
	}

	repo.EXPECT().GetSensorById(client, gomock.Any(), id).Return(nil, apperror.ErrEntityNotFound)
	repo.EXPECT().CreateSensor(client, gomock.Any(), sensor).Return(nil)
	err := service.CreateSensor(client, context.Background(), sensor)
	assert.NoError(t, err)
}

func TestService_UpdateSensorAnalyzer(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	sId := "20081UC-023"
	anId := "20081UC-061"
	repo.EXPECT().UpdateSensorAnalyzer(client, gomock.Any(), sId, anId).Return(nil)
	err := service.UpdateSensorAnalyzer(client, context.Background(), sId, anId)
	assert.NoError(t, err)
}

func TestService_DeleteSensor(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := "20081UC-061"
	repo.EXPECT().DeleteSensor(client, gomock.Any(), id).Return(nil)
	err := service.DeleteSensor(client, context.Background(), id)
	assert.NoError(t, err)
}
