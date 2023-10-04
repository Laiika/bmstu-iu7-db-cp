package service

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"errors"
	"strings"
)

type IGasRepo interface {
	GetGasByName(client Client, ctx context.Context, name string) (*entity.Gas, error)
	GetAllGases(client Client, ctx context.Context) (entity.Gases, error)
	GetTypeGases(client Client, ctx context.Context, anType string) (entity.Gases, error)
	CreateGas(client Client, ctx context.Context, gas *entity.Gas) error
}

func (s *Service) GetAllGases(client Client, ctx context.Context) (entity.Gases, error) {
	gases, err := s.repo.GetAllGases(client, ctx)
	if err != nil {
		s.log.Error(err)
	}

	return gases, err
}

func (s *Service) GetTypeGases(client Client, ctx context.Context, anType string) (entity.Gases, error) {
	anType = strings.ToUpper(anType)
	gases, err := s.repo.GetTypeGases(client, ctx, anType)
	if err != nil {
		s.log.Error(err)
	}

	return gases, err
}

func (s *Service) CreateGas(client Client, ctx context.Context, gas *entity.Gas) error {
	gas.Name = strings.ToLower(gas.Name)
	if err := gas.IsValid(); err != nil {
		s.log.Error(err)
		return err
	}

	_, err := s.repo.GetGasByName(client, ctx, gas.Name)
	if err == nil {
		s.log.Error(apperror.ErrEntityExists)
		return apperror.ErrEntityExists
	} else if !errors.Is(err, apperror.ErrEntityNotFound) {
		s.log.Error(err)
		return err
	}

	err = s.repo.CreateGas(client, ctx, gas)
	if err != nil {
		s.log.Error(err)
	}

	return err
}
