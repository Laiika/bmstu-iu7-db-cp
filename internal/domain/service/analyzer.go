package service

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"github.com/pkg/errors"
	"strings"
)

type IAnalyzerRepo interface {
	GetAnalyzerById(client Client, ctx context.Context, id string) (*entity.Analyzer, error)
	GetAllAnalyzers(client Client, ctx context.Context) (entity.Analyzers, error)
	GetTypeAnalyzers(client Client, ctx context.Context, anType string) (entity.Analyzers, error)
	CreateAnalyzer(client Client, ctx context.Context, analyzer *entity.Analyzer) error
	DeleteAnalyzer(client Client, ctx context.Context, id string) error
}

func (s *Service) GetAnalyzerById(client Client, ctx context.Context, id string) (*entity.Analyzer, error) {
	analyzer, err := s.repo.GetAnalyzerById(client, ctx, id)
	if err != nil {
		s.log.Error(err)
	}

	return analyzer, err
}

func (s *Service) GetAllAnalyzers(client Client, ctx context.Context) (entity.Analyzers, error) {
	analyzers, err := s.repo.GetAllAnalyzers(client, ctx)
	if err != nil {
		s.log.Error(err)
	}

	return analyzers, err
}

func (s *Service) GetTypeAnalyzers(client Client, ctx context.Context, anType string) (entity.Analyzers, error) {
	anType = strings.ToUpper(anType)
	analyzers, err := s.repo.GetTypeAnalyzers(client, ctx, anType)
	if err != nil {
		s.log.Error(err)
	}

	return analyzers, err
}

func (s *Service) CreateAnalyzer(client Client, ctx context.Context, analyzer *entity.Analyzer) error {
	if err := analyzer.IsValid(); err != nil {
		s.log.Error(err)
		return err
	}

	_, err := s.repo.GetAnalyzerById(client, ctx, analyzer.Id)
	if err == nil {
		s.log.Error(apperror.ErrEntityExists)
		return apperror.ErrEntityExists
	} else if !errors.Is(err, apperror.ErrEntityNotFound) {
		s.log.Error(err)
		return err
	}

	analyzer.Type = strings.ToUpper(analyzer.Type)
	err = s.repo.CreateAnalyzer(client, ctx, analyzer)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *Service) DeleteAnalyzer(client Client, ctx context.Context, id string) error {
	err := s.repo.DeleteAnalyzer(client, ctx, id)
	if err != nil {
		s.log.Error(err)
	}

	return err
}
