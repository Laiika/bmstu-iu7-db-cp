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

type SensorRepo struct {
}

func NewSensorRepo() *SensorRepo {
	return &SensorRepo{}
}

func (r *AnalyzerRepo) GetSensorById(client service.Client, ctx context.Context, id string) (*entity.Sensor, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, type, gas, analyzer_id, low_limit_alarm, upper_limit_alarm
		FROM sensors
		WHERE id = $1
	`
	var an entity.Sensor
	err := pgClient.QueryRow(ctx, q, id).Scan(&an.Id, &an.Type, &an.Gas, &an.AnalyzerId, &an.LowLimitAlarm, &an.UpperLimitAlarm)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperror.ErrEntityNotFound
		}
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return &an, nil
}

func (r *SensorRepo) GetAllSensors(client service.Client, ctx context.Context) (entity.Sensors, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, type, gas, analyzer_id, low_limit_alarm, upper_limit_alarm
		FROM sensors
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	sensors := make(entity.Sensors, 0)
	for rows.Next() {
		var an entity.Sensor

		err = rows.Scan(&an.Id, &an.Type, &an.Gas, &an.AnalyzerId, &an.LowLimitAlarm, &an.UpperLimitAlarm)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		sensors = append(sensors, &an)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return sensors, nil
}

func (r *SensorRepo) GetAnalyzerSensors(client service.Client, ctx context.Context, anId string) (entity.Sensors, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, type, gas, analyzer_id, low_limit_alarm, upper_limit_alarm
		FROM sensors
		WHERE analyzer_id = $1
	`
	rows, err := pgClient.Query(ctx, q, anId)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	sensors := make(entity.Sensors, 0)
	for rows.Next() {
		var an entity.Sensor

		err = rows.Scan(&an.Id, &an.Type, &an.Gas, &an.AnalyzerId, &an.LowLimitAlarm, &an.UpperLimitAlarm)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		sensors = append(sensors, &an)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return sensors, nil
}

func (r *SensorRepo) CreateSensor(client service.Client, ctx context.Context, sensor *entity.Sensor) error {
	pgClient := client.(postgresql.Client)
	q := `
		INSERT INTO sensors
		    (id, type, gas, analyzer_id, low_limit_alarm, upper_limit_alarm) 
		VALUES 
		    ($1, $2, $3, $4, $5, $6)
	`
	_, err := pgClient.Exec(ctx, q, sensor.Id, sensor.Type, sensor.Gas, sensor.AnalyzerId, sensor.LowLimitAlarm, sensor.UpperLimitAlarm)
	if err != nil {
		return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return nil
}

func (r *SensorRepo) UpdateSensorAnalyzer(client service.Client, ctx context.Context, sId string, anId string) error {
	pgClient := client.(postgresql.Client)
	q := `
		UPDATE sensors
		SET 
		    analyzer_id = $1
		WHERE id = $2
	`
	commandTag, err := pgClient.Exec(ctx, q, anId, sId)
	if err != nil {
		return pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return apperror.ErrEntityNotFound
	}

	return nil
}

func (r *SensorRepo) DeleteSensor(client service.Client, ctx context.Context, id string) error {
	pgClient := client.(postgresql.Client)
	q := `
		DELETE FROM sensors
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
