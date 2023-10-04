package postgres

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"db_cp_6_sem/internal/domain/service"
	"db_cp_6_sem/pkg/client/postgresql"
	"errors"
	"github.com/jackc/pgx/v5"
	pkgErrors "github.com/pkg/errors"
)

type AnalyzerRepo struct {
}

func NewAnalyzerRepo() *AnalyzerRepo {
	return &AnalyzerRepo{}
}

func (r *AnalyzerRepo) GetAnalyzerById(client service.Client, ctx context.Context, id string) (*entity.Analyzer, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, type, part_number, job_number, software_version
		FROM gas_analyzers
		WHERE id = $1
	`
	var an entity.Analyzer
	err := pgClient.QueryRow(ctx, q, id).Scan(&an.Id, &an.Type, &an.PartNumber, &an.JobNumber, &an.SoftwareVersion)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrEntityNotFound
		}
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return &an, nil
}

func (r *AnalyzerRepo) GetAllAnalyzers(client service.Client, ctx context.Context) (entity.Analyzers, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, type, part_number, job_number, software_version
		FROM gas_analyzers
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	analyzers := make(entity.Analyzers, 0)
	for rows.Next() {
		var an entity.Analyzer

		err = rows.Scan(&an.Id, &an.Type, &an.PartNumber, &an.JobNumber, &an.SoftwareVersion)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		analyzers = append(analyzers, &an)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return analyzers, nil
}

func (r *AnalyzerRepo) GetTypeAnalyzers(client service.Client, ctx context.Context, anType string) (entity.Analyzers, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, type, part_number, job_number, software_version
		FROM gas_analyzers
		WHERE type = $1
	`
	rows, err := pgClient.Query(ctx, q, anType)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	analyzers := make(entity.Analyzers, 0)
	for rows.Next() {
		var an entity.Analyzer

		err = rows.Scan(&an.Id, &an.Type, &an.PartNumber, &an.JobNumber, &an.SoftwareVersion)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		analyzers = append(analyzers, &an)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return analyzers, nil
}

func (r *AnalyzerRepo) CreateAnalyzer(client service.Client, ctx context.Context, analyzer *entity.Analyzer) error {
	pgClient := client.(postgresql.Client)
	q := `
		INSERT INTO gas_analyzers
		    (id, type, part_number, job_number, software_version) 
		VALUES 
		    ($1, $2, $3, $4, $5)
	`
	_, err := pgClient.Exec(ctx, q, analyzer.Id, analyzer.Type, analyzer.PartNumber, analyzer.JobNumber, analyzer.SoftwareVersion)
	if err != nil {
		return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return nil
}

func (r *AnalyzerRepo) DeleteAnalyzer(client service.Client, ctx context.Context, id string) error {
	pgClient := client.(postgresql.Client)
	q := `
		DELETE FROM gas_analyzers
		WHERE id = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, id)
	if err != nil {
		return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return apperror.ErrEntityNotFound
	}

	return nil
}
