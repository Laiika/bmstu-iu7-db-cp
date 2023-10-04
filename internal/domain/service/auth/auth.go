package auth

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"db_cp_6_sem/internal/domain/service"
	"db_cp_6_sem/pkg/logger"
	"errors"
	pkgErrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type AuthService struct {
	user     service.Client
	empl     service.Client
	admin    service.Client
	mx       sync.RWMutex
	sessions map[string]*session
	log      *logger.Logger
}

func NewAuthService(user service.Client, empl service.Client, admin service.Client, log *logger.Logger) *AuthService {
	return &AuthService{
		user:     user,
		empl:     empl,
		admin:    admin,
		mx:       sync.RWMutex{},
		sessions: make(map[string]*session),
		log:      log,
	}
}

func (s *AuthService) Login(ctx context.Context, service *service.Service, data *entity.Auth) (string, error) {
	curUser, err := service.GetUserByName(s.admin, ctx, data.Username)
	if err != nil {
		if errors.Is(err, apperror.ErrEntityNotFound) {
			s.log.Error(pkgErrors.WithMessage(apperror.ErrUnauthorized, "no user with that name"))
			return "", pkgErrors.WithMessage(apperror.ErrUnauthorized, "no user with that name")
		}
		s.log.Error(pkgErrors.WithMessage(apperror.ErrInternal, err.Error()))
		return "", pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}
	if !checkPassword(data.Password, curUser.Password) {
		s.log.Error(pkgErrors.WithMessage(apperror.ErrUnauthorized, "incorrect password"))
		return "", pkgErrors.WithMessage(apperror.ErrUnauthorized, "incorrect password")
	}

	ses := NewSession(s.user, s.empl, s.admin, curUser.Id, curUser.Role)

	s.mx.Lock()
	s.sessions[ses.GetToken()] = ses
	s.mx.Unlock()
	s.log.Infof("add session: token %s, role %s", ses.GetToken(), ses.GetRole())

	return ses.GetToken(), nil
}

func (s *AuthService) Logout(token string) error {
	s.mx.RLock()
	ses, ok := s.sessions[token]
	s.mx.RUnlock()
	if !ok {
		s.log.Error(pkgErrors.WithMessage(apperror.ErrSessionNotExists, token))
		return pkgErrors.WithMessage(apperror.ErrSessionNotExists, token)
	}

	s.log.Infof("remove session: token %s, role %s", ses.GetToken(), ses.GetRole())
	s.mx.Lock()
	delete(s.sessions, token)
	s.mx.Unlock()

	return nil
}

func (s *AuthService) GetSession(token string) bool {
	s.mx.RLock()
	_, ok := s.sessions[token]
	s.mx.RUnlock()

	return ok
}

func (s *AuthService) GetClient(token string) (service.Client, error) {
	s.mx.RLock()
	ses, ok := s.sessions[token]
	s.mx.RUnlock()
	if !ok {
		s.log.Error(pkgErrors.WithMessage(apperror.ErrSessionNotExists, token))
		return nil, pkgErrors.WithMessage(apperror.ErrSessionNotExists, token)
	}

	return ses.GetClient(), nil
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
