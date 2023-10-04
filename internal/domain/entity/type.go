package entity

import (
	"db_cp_6_sem/internal/apperror"
)

type Type struct {
	Name       string `json:"name"`
	MaxSensors int    `json:"max_sensors"`
}

type Types []*Type

type CreateType struct {
	Name       string   `json:"name"`
	MaxSensors int      `json:"max_sensors"`
	Gases      []string `json:"gases"`
}

func (t *CreateType) IsValid() error {
	var err error

	switch {
	case t.Name == "":
		err = apperror.ErrInvalidName
	case t.MaxSensors <= 0:
		err = apperror.ErrInvalidSensorsNumber
	}

	return err
}
