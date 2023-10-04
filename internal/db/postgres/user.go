package postgres

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"db_cp_6_sem/internal/domain/service"
	"db_cp_6_sem/pkg/client/postgresql"
	"github.com/jackc/pgx/v5"
	pkgErrors "github.com/pkg/errors"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUserById(client service.Client, ctx context.Context, id int) (*entity.User, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, name, password, role
		FROM users
		WHERE id = $1
	`
	var an entity.User
	err := pgClient.QueryRow(ctx, q, id).Scan(&an.Id, &an.Name, &an.Password, &an.Role)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrEntityNotFound
		}
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return &an, nil
}

func (r *UserRepo) GetUserByName(client service.Client, ctx context.Context, name string) (*entity.User, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, name, password, role
		FROM users
		WHERE name = $1
	`
	var an entity.User
	err := pgClient.QueryRow(ctx, q, name).Scan(&an.Id, &an.Name, &an.Password, &an.Role)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrEntityNotFound
		}
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return &an, nil
}

func (r *UserRepo) GetAllUsers(client service.Client, ctx context.Context) (entity.Users, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, name, password, role
		FROM users
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	users := make(entity.Users, 0)
	for rows.Next() {
		var an entity.User

		err = rows.Scan(&an.Id, &an.Name, &an.Password, &an.Role)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		users = append(users, &an)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return users, nil
}

func (r *UserRepo) CreateUser(client service.Client, ctx context.Context, user *entity.User) (int, error) {
	pgClient := client.(postgresql.Client)
	q := `
		INSERT INTO users
		    (name, password, role) 
		VALUES 
		    ($1, $2, $3) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, user.Name, user.Password, user.Role).Scan(&id)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return id, nil
}

func (r *UserRepo) UpdateUserRole(client service.Client, ctx context.Context, id int, role string) error {
	pgClient := client.(postgresql.Client)
	q := `
		UPDATE users
		SET 
		    role = $1
		WHERE id = $2
	`
	commandTag, err := pgClient.Exec(ctx, q, role, id)
	if err != nil {
		return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return apperror.ErrEntityNotFound
	}

	return nil
}

func (r *UserRepo) DeleteUser(client service.Client, ctx context.Context, id int) error {
	pgClient := client.(postgresql.Client)
	q := `
		DELETE FROM users
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
