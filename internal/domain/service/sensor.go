package service

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"errors"
	"strings"
)

type ISensorRepo interface {
	GetSensorById(client Client, ctx context.Context, id string) (*entity.Sensor, error)
	GetAllSensors(client Client, ctx context.Context) (entity.Sensors, error)
	GetAnalyzerSensors(client Client, ctx context.Context, anId string) (entity.Sensors, error)
	CreateSensor(client Client, ctx context.Context, sensor *entity.Sensor) error
	UpdateSensorAnalyzer(client Client, ctx context.Context, sId string, anId string) error
	DeleteSensor(client Client, ctx context.Context, id string) error
}

func (s *Service) GetAllSensors(client Client, ctx context.Context) (entity.Sensors, error) {
	sensors, err := s.repo.GetAllSensors(client, ctx)
	if err != nil {
		s.log.Error(err)
	}

	return sensors, err
}

func (s *Service) GetAnalyzerSensors(client Client, ctx context.Context, anId string) (entity.Sensors, error) {
	sensors, err := s.repo.GetAnalyzerSensors(client, ctx, anId)
	if err != nil {
		s.log.Error(err)
	}

	return sensors, err
}

func (s *Service) CreateSensor(client Client, ctx context.Context, sensor *entity.Sensor) error {
	if err := sensor.IsValid(); err != nil {
		s.log.Error(err)
		return err
	}

	_, err := s.repo.GetSensorById(client, ctx, sensor.Id)
	if err == nil {
		s.log.Error(apperror.ErrEntityExists)
		return apperror.ErrEntityExists
	} else if !errors.Is(err, apperror.ErrEntityNotFound) {
		s.log.Error(err)
		return err
	}

	sensor.Gas = strings.ToLower(sensor.Gas)
	err = s.repo.CreateSensor(client, ctx, sensor)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *Service) UpdateSensorAnalyzer(client Client, ctx context.Context, sId string, anId string) error {
	err := s.repo.UpdateSensorAnalyzer(client, ctx, sId, anId)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *Service) DeleteSensor(client Client, ctx context.Context, id string) error {
	err := s.repo.DeleteSensor(client, ctx, id)
	if err != nil {
		s.log.Error(err)
	}

	return err
}
