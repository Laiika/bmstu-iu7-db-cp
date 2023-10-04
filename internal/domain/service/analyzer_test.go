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

func TestService_GetAnalyzerById(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := "20081UC-061"
	expectedAnalyzer := &entity.Analyzer{
		Id:              id,
		Type:            "VENTIS MX4",
		PartNumber:      "VTS-L0001100709",
		JobNumber:       "20091U",
		SoftwareVersion: "10.11.01",
	}

	repo.EXPECT().GetAnalyzerById(client, gomock.Any(), id).Return(expectedAnalyzer, nil)
	analyzer, err := service.GetAnalyzerById(client, context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnalyzer, analyzer)
}

func TestService_GetAllAnalyzers(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	expectedAnalyzers := entity.Analyzers{
		&entity.Analyzer{
			Id:              "20081UC-061",
			Type:            "VENTIS MX4",
			PartNumber:      "VTS-L0001100709",
			JobNumber:       "20091U",
			SoftwareVersion: "10.11.01",
		},
		&entity.Analyzer{
			Id:              "20081UC-023",
			Type:            "VENTIS MX4",
			PartNumber:      "VTS-L0001104567",
			JobNumber:       "20121U",
			SoftwareVersion: "4.11.01",
		},
	}

	repo.EXPECT().GetAllAnalyzers(client, gomock.Any()).Return(expectedAnalyzers, nil)
	analyzers, err := service.GetAllAnalyzers(client, context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedAnalyzers, analyzers)
}

func TestService_GetTypeAnalyzers(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	anType := "VENTIS MX4"
	expectedAnalyzers := entity.Analyzers{
		&entity.Analyzer{
			Id:              "20081UC-061",
			Type:            anType,
			PartNumber:      "VTS-L0001100709",
			JobNumber:       "20091U",
			SoftwareVersion: "10.11.01",
		},
		&entity.Analyzer{
			Id:              "20081UC-023",
			Type:            anType,
			PartNumber:      "VTS-L0001104567",
			JobNumber:       "20121U",
			SoftwareVersion: "4.11.01",
		},
	}

	repo.EXPECT().GetTypeAnalyzers(client, gomock.Any(), anType).Return(expectedAnalyzers, nil)
	analyzers, err := service.GetTypeAnalyzers(client, context.Background(), anType)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnalyzers, analyzers)
}

func TestService_CreateAnalyzer(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := "20081UC-023"
	analyzer := &entity.Analyzer{
		Id:              id,
		Type:            "VENTIS MX4",
		PartNumber:      "VTS-L0001104567",
		JobNumber:       "20121U",
		SoftwareVersion: "4.11.01",
	}

	repo.EXPECT().GetAnalyzerById(client, gomock.Any(), id).Return(nil, apperror.ErrEntityNotFound)
	repo.EXPECT().CreateAnalyzer(client, gomock.Any(), analyzer).Return(nil)
	err := service.CreateAnalyzer(client, context.Background(), analyzer)
	assert.NoError(t, err)
}

func TestService_DeleteAnalyzer(t *testing.T) {
	log := logger.GetLogger()
	var client Client

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repo := NewMockIRepo(ctl)
	service := NewService(repo, log)

	id := "20081UC-061"
	repo.EXPECT().DeleteAnalyzer(client, gomock.Any(), id).Return(nil)
	err := service.DeleteAnalyzer(client, context.Background(), id)
	assert.NoError(t, err)
}
