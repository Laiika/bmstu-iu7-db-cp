package service

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"errors"
	pkgErrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type IUserRepo interface {
	GetUserById(client Client, ctx context.Context, id int) (*entity.User, error)
	GetUserByName(client Client, ctx context.Context, name string) (*entity.User, error)
	GetAllUsers(client Client, ctx context.Context) (entity.Users, error)
	CreateUser(client Client, ctx context.Context, user *entity.User) (int, error)
	UpdateUserRole(client Client, ctx context.Context, id int, role string) error
	DeleteUser(client Client, ctx context.Context, id int) error
}

func (s *Service) GetUserById(client Client, ctx context.Context, id int) (*entity.User, error) {
	user, err := s.repo.GetUserById(client, ctx, id)
	if err != nil {
		s.log.Error(err)
	}

	return user, err
}

func (s *Service) GetUserByName(client Client, ctx context.Context, name string) (*entity.User, error) {
	name = strings.ToLower(name)
	user, err := s.repo.GetUserByName(client, ctx, name)
	if err != nil {
		s.log.Error(err)
	}

	return user, err
}

func (s *Service) GetAllUsers(client Client, ctx context.Context) (entity.Users, error) {
	users, err := s.repo.GetAllUsers(client, ctx)
	if err != nil {
		s.log.Error(err)
	}

	return users, err
}

func (s *Service) CreateUser(client Client, ctx context.Context, user *entity.CreateUser) (int, error) {
	user.Role = strings.ToLower(user.Role)
	if err := user.IsValid(); err != nil {
		s.log.Error(err)
		return 0, err
	}

	_, err := s.repo.GetUserByName(client, ctx, user.Name)
	if err == nil {
		s.log.Error(apperror.ErrEntityExists)
		return 0, apperror.ErrEntityExists
	} else if !errors.Is(err, apperror.ErrEntityNotFound) {
		s.log.Error(err)
		return 0, err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		s.log.Error(pkgErrors.WithMessage(apperror.ErrInternal, err.Error()))
		return 0, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	us := &entity.User{
		Name:     user.Name,
		Password: string(bytes),
		Role:     user.Role,
	}
	id, err := s.repo.CreateUser(client, ctx, us)
	if err != nil {
		s.log.Error(err)
	}

	return id, err
}

func (s *Service) UpdateUserRole(client Client, ctx context.Context, id int, role string) error {
	role = strings.ToLower(role)
	if role != "admin" && role != "employee" && role != "user" {
		s.log.Error(apperror.ErrInvalidRole)
		return apperror.ErrInvalidRole
	}

	err := s.repo.UpdateUserRole(client, ctx, id, role)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *Service) DeleteUser(client Client, ctx context.Context, id int) error {
	err := s.repo.DeleteUser(client, ctx, id)
	if err != nil {
		s.log.Error(err)
	}

	return err
}
