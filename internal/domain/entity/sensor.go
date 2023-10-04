package entity

import (
	"db_cp_6_sem/internal/apperror"
)

type Sensor struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	AnalyzerId      string `json:"analyzer_id"`
	Gas             string `json:"gas"`
	LowLimitAlarm   string `json:"low_limit_alarm"`
	UpperLimitAlarm string `json:"upper_limit_alarm"`
}

type Sensors []*Sensor

func (s *Sensor) IsValid() error {
	var err error

	switch {
	case s.Id == "":
		err = apperror.ErrInvalidId
	case s.Type == "":
		err = apperror.ErrInvalidType
	case s.Gas == "":
		err = apperror.ErrInvalidGas
	}

	return err
}
