package entity

import (
	"db_cp_6_sem/internal/apperror"
)

type Analyzer struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	PartNumber      string `json:"part_number"`
	JobNumber       string `json:"job_number"`
	SoftwareVersion string `json:"software_version"`
}

type Analyzers []*Analyzer

func (an *Analyzer) IsValid() error {
	var err error

	switch {
	case an.Id == "":
		err = apperror.ErrInvalidId
	case an.Type == "":
		err = apperror.ErrInvalidType
	case an.SoftwareVersion == "":
		err = apperror.ErrInvalidSoftware
	}

	return err
}
