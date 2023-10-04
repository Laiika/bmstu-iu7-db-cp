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

type TypeRepo struct {
}

func NewTypeRepo() *TypeRepo {
	return &TypeRepo{}
}

func (r *TypeRepo) GetTypeByName(client service.Client, ctx context.Context, name string) (*entity.Type, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT name, max_sensors
		FROM analyzer_types
		WHERE name = $1
	`
	var g entity.Type
	err := pgClient.QueryRow(ctx, q, name).Scan(&g.Name, &g.MaxSensors)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrEntityNotFound
		}
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return &g, nil
}

func (r *TypeRepo) GetAllTypes(client service.Client, ctx context.Context) (entity.Types, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT name, max_sensors
		FROM analyzer_types
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	types := make(entity.Types, 0)
	for rows.Next() {
		var g entity.Type

		err = rows.Scan(&g.Name, &g.MaxSensors)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		types = append(types, &g)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return types, nil
}

func (r *TypeRepo) CreateType(client service.Client, ctx context.Context, anType *entity.Type, gases []string) error {
	pgClient := client.(postgresql.Client)
	q := `
		INSERT INTO analyzer_types
		    (name, max_sensors) 
		VALUES 
		    ($1::text, $2::int)
	`
	_, err := pgClient.Exec(ctx, q, anType.Name, anType.MaxSensors)
	if err != nil {
		return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	for _, gas := range gases {
		q = `
		INSERT INTO types_gases
		    (gas, analyzer_type) 
		VALUES 
		    ($1::text, $2::text)
		`
		_, err = pgClient.Exec(ctx, q, gas, anType.Name)
		if err != nil {
			return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}
	}

	return nil
}

func (r *TypeRepo) DeleteType(client service.Client, ctx context.Context, name string) error {
	pgClient := client.(postgresql.Client)
	q := `
		DELETE FROM analyzer_types
		WHERE name = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, name)
	if err != nil {
		return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return apperror.ErrEntityNotFound
	}

	return nil
}
