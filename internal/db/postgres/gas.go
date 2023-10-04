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

type GasRepo struct {
}

func NewGasRepo() *GasRepo {
	return &GasRepo{}
}

func (r *GasRepo) GetGasByName(client service.Client, ctx context.Context, name string) (*entity.Gas, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT name, formula, type
		FROM gases
		WHERE name = $1
	`
	var g entity.Gas
	err := pgClient.QueryRow(ctx, q, name).Scan(&g.Name, &g.Formula, &g.Type)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrEntityNotFound
		}
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return &g, nil
}

func (r *GasRepo) GetAllGases(client service.Client, ctx context.Context) (entity.Gases, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT name, formula, type
		FROM gases
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	gases := make(entity.Gases, 0)
	for rows.Next() {
		var g entity.Gas

		err = rows.Scan(&g.Name, &g.Formula, &g.Type)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		gases = append(gases, &g)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return gases, nil
}

func (r *GasRepo) GetTypeGases(client service.Client, ctx context.Context, anType string) (entity.Gases, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT g.name, g.formula, g.type
		FROM gases g
		JOIN (SELECT gas
			  FROM types_gases
			  WHERE analyzer_type = $1) AS tg
		ON tg.gas = g.name
	`
	rows, err := pgClient.Query(ctx, q, anType)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	gases := make(entity.Gases, 0)
	for rows.Next() {
		var g entity.Gas

		err = rows.Scan(&g.Name, &g.Formula, &g.Type)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		gases = append(gases, &g)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return gases, nil
}

func (r *GasRepo) CreateGas(client service.Client, ctx context.Context, gas *entity.Gas) error {
	pgClient := client.(postgresql.Client)
	q := `
		INSERT INTO gases
		    (name, formula, type) 
		VALUES 
		    ($1, $2, $3)
	`
	_, err := pgClient.Exec(ctx, q, gas.Name, gas.Formula, gas.Type)
	if err != nil {
		return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return nil
}
