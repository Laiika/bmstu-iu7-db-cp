package apperror

import "errors"

var (
	ErrInvalidId   = errors.New("invalid serial number")
	ErrInvalidType = errors.New("invalid type")
	ErrInvalidName = errors.New("invalid name")

	ErrInvalidSoftware      = errors.New("invalid software version")
	ErrInvalidLimits        = errors.New("the lower limit exceeds the upper")
	ErrInvalidSensorsNumber = errors.New("invalid max number of sensors")
	ErrInvalidSensor        = errors.New("invalid sensor id")
	ErrInvalidGas           = errors.New("invalid gas name")

	ErrInvalidUserName = errors.New("invalid user name")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidRole     = errors.New("invalid role")

	ErrUnauthorized     = errors.New("unauthorized access")
	ErrSessionNotExists = errors.New("user session not found")

	ErrEntityExists   = errors.New("this entity already exists in database")
	ErrInternal       = errors.New("internal server error")
	ErrEntityNotFound = errors.New("entity not found")
)
