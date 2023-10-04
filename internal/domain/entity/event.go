package entity

import (
	"db_cp_6_sem/internal/apperror"
	"time"
)

type Event struct {
	Id           int       `json:"id"`
	SignalTime   time.Time `json:"signal_time"`
	SensorId     string    `json:"sensor_id"`
	PeakReadings float32   `json:"peak_readings"`
}

type Events []*Event

type CreateEvent struct {
	SignalTime   string  `json:"signal_time"`
	SensorId     string  `json:"sensor_id"`
	PeakReadings float32 `json:"peak_readings"`
}

func (ev *CreateEvent) IsValid() error {
	var err error

	if ev.SensorId == "" {
		err = apperror.ErrInvalidSensor
	}

	return err
}
