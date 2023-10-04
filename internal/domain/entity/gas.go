package entity

import "db_cp_6_sem/internal/apperror"

type Gas struct {
	Name    string `json:"name"`
	Formula string `json:"formula"`
	Type    string `json:"type"`
}

type Gases []*Gas

func (gas *Gas) IsValid() error {
	var err error

	switch {
	case gas.Name == "":
		err = apperror.ErrInvalidName
	case gas.Type == "":
		err = apperror.ErrInvalidType
	}

	return err
}
