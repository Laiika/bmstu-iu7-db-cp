package service

import (
	"context"
	"db_cp_6_sem/internal/domain/entity"
	"db_cp_6_sem/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_GetSensorEvents(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	sId := "14032ER395"
	expectedEvents := entity.Events{
		&entity.Event{
			Id:           1,
			SignalTime:   time.Now(),
			SensorId:     sId,
			PeakReadings: 11.1,
		},
		&entity.Event{
			Id:           2,
			SignalTime:   time.Now(),
			SensorId:     sId,
			PeakReadings: 21.1,
		},
	}

	repo.EXPECT().GetSensorEvents(client, gomock.Any(), sId).Return(expectedEvents, nil)
	events, err := service.GetSensorEvents(client, context.Background(), sId)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, events)
}

func TestService_GetEventsBySignalTime(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	left := time.Now()
	right := left.Add(12 * time.Minute)
	expectedEvents := entity.Events{
		&entity.Event{
			Id:           1,
			SignalTime:   left.Add(5 * time.Second),
			SensorId:     "14032ER395",
			PeakReadings: 11.1,
		},
		&entity.Event{
			Id:           2,
			SignalTime:   left.Add(10 * time.Minute),
			SensorId:     "22101BL013",
			PeakReadings: 21.1,
		},
	}

	repo.EXPECT().GetEventsBySignalTime(client, gomock.Any(), left, right).Return(expectedEvents, nil)
	events, err := service.GetEventsBySignalTime(client, context.Background(), left, right)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, events)
}

func TestService_CreateEvent(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	tm := "2020-02-04"
	signalTime, _ := time.Parse("2006-01-02", tm)
	event := &entity.Event{
		SignalTime:   signalTime,
		SensorId:     "22101BL013",
		PeakReadings: 21.1,
	}
	ev := &entity.CreateEvent{
		SignalTime:   tm,
		SensorId:     "22101BL013",
		PeakReadings: 21.1,
	}

	id := 1
	repo.EXPECT().CreateEvent(client, gomock.Any(), event).Return(id, nil)
	id2, err := service.CreateEvent(client, context.Background(), ev)
	assert.NoError(t, err)
	assert.Equal(t, id, id2)
}

func TestService_DeleteEvent(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := 1
	repo.EXPECT().DeleteEvent(client, gomock.Any(), id).Return(nil)
	err := service.DeleteEvent(client, context.Background(), id)
	assert.NoError(t, err)
}
