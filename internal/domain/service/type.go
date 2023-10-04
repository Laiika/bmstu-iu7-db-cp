package service

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/pkg/errors"
	"strings"
)

type ITypeRepo interface {
	GetTypeByName(client Client, ctx context.Context, name string) (*entity.Type, error)
	GetAllTypes(client Client, ctx context.Context) (entity.Types, error)
	CreateType(client Client, ctx context.Context, anType *entity.Type, gases []string) error
	DeleteType(client Client, ctx context.Context, id string) error
}

func (s *Service) GetAllTypes(client Client, ctx context.Context) (entity.Types, error) {
	types, err := s.repo.GetAllTypes(client, ctx)
	if err != nil {
		s.log.Error(err)
	}

	return types, err
}

func (s *Service) CreateType(client Client, ctx context.Context, anType *entity.CreateType) error {
	if err := anType.IsValid(); err != nil {
		s.log.Error(err)
		return err
	}

	name := strings.ToUpper(anType.Name)
	_, err := s.repo.GetTypeByName(client, ctx, name)
	if err == nil {
		s.log.Error(apperror.ErrEntityExists)
		return apperror.ErrEntityExists
	} else if !errors.Is(err, apperror.ErrEntityNotFound) {
		s.log.Error(err)
		return err
	}

	t := &entity.Type{
		Name:       name,
		MaxSensors: anType.MaxSensors,
	}
	err = s.repo.CreateType(client, ctx, t, anType.Gases)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *Service) DeleteType(client Client, ctx context.Context, id string) error {
	id = strings.ToUpper(id)
	err := s.repo.DeleteType(client, ctx, id)
	if err != nil {
		s.log.Error(err)
	}

	return err
}
