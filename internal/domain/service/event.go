package service

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"time"
)

type IEventRepo interface {
	GetEventsBySignalTime(client Client, ctx context.Context, left time.Time, right time.Time) (entity.Events, error)
	GetSensorEvents(client Client, ctx context.Context, sId string) (entity.Events, error)
	CreateEvent(client Client, ctx context.Context, event *entity.Event) (int, error)
	DeleteEvent(client Client, ctx context.Context, id int) error
}

func (s *Service) GetEventsBySignalTime(client Client, ctx context.Context, left time.Time, right time.Time) (entity.Events, error) {
	if left.After(right) {
		s.log.Error(apperror.ErrInvalidLimits)
		return nil, apperror.ErrInvalidLimits
	}

	events, err := s.repo.GetEventsBySignalTime(client, ctx, left, right)
	if err != nil {
		s.log.Error(err)
	}

	return events, err
}

func (s *Service) GetSensorEvents(client Client, ctx context.Context, sId string) (entity.Events, error) {
	events, err := s.repo.GetSensorEvents(client, ctx, sId)
	if err != nil {
		s.log.Error(err)
	}

	return events, err
}

func (s *Service) CreateEvent(client Client, ctx context.Context, event *entity.CreateEvent) (int, error) {
	if err := event.IsValid(); err != nil {
		s.log.Error(err)
		return 0, err
	}

	signalTime, _ := time.Parse("2006-01-02 03:04:05", event.SignalTime)
	ev := &entity.Event{
		SignalTime:   signalTime,
		SensorId:     event.SensorId,
		PeakReadings: event.PeakReadings,
	}
	id, err := s.repo.CreateEvent(client, ctx, ev)
	if err != nil {
		s.log.Error(err)
	}

	return id, err
}

func (s *Service) DeleteEvent(client Client, ctx context.Context, id int) error {
	err := s.repo.DeleteEvent(client, ctx, id)
	if err != nil {
		s.log.Error(err)
	}

	return err
}
