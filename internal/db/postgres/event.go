package postgres

import (
	"context"
	"db_cp_6_sem/internal/apperror"
	"db_cp_6_sem/internal/domain/entity"
	"db_cp_6_sem/internal/domain/service"
	"db_cp_6_sem/pkg/client/postgresql"
	pkgErrors "github.com/pkg/errors"
	"time"
)

type EventRepo struct {
}

func NewEventRepo() *EventRepo {
	return &EventRepo{}
}

func (r *EventRepo) GetEventsBySignalTime(client service.Client, ctx context.Context, left time.Time, right time.Time) (entity.Events, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, signal_time, sensor_id, peak_readings
		FROM events
		WHERE signal_time >= $1::timestamp AND signal_time <= $2::timestamp
	`
	rows, err := pgClient.Query(ctx, q, left.Format("2006-01-02 15:04:05"), right.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	events := make(entity.Events, 0)
	for rows.Next() {
		var ev entity.Event

		err = rows.Scan(&ev.Id, &ev.SignalTime, &ev.SensorId, &ev.PeakReadings)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		events = append(events, &ev)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return events, nil
}

func (r *EventRepo) GetSensorEvents(client service.Client, ctx context.Context, sId string) (entity.Events, error) {
	pgClient := client.(postgresql.Client)
	q := `
		SELECT id, signal_time, sensor_id, peak_readings
		FROM events
		WHERE sensor_id = $1
	`
	rows, err := pgClient.Query(ctx, q, sId)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	events := make(entity.Events, 0)
	for rows.Next() {
		var ev entity.Event

		err = rows.Scan(&ev.Id, &ev.SignalTime, &ev.SensorId, &ev.PeakReadings)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
		}

		events = append(events, &ev)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return events, nil
}

func (r *EventRepo) CreateEvent(client service.Client, ctx context.Context, event *entity.Event) (int, error) {
	pgClient := client.(postgresql.Client)
	q := `
		INSERT INTO events
		    (signal_time, sensor_id, peak_readings) 
		VALUES 
		    ($1, $2, $3) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, event.SignalTime, event.SensorId, event.PeakReadings).Scan(&id)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperror.ErrInternal, err.Error())
	}

	return id, nil
}

func (r *EventRepo) DeleteEvent(client service.Client, ctx context.Context, id int) error {
	pgClient := client.(postgresql.Client)
	q := `
		DELETE FROM events
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
