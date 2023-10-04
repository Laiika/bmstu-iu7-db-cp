package entity

import "db_cp_6_sem/internal/apperror"

type User struct {
	Id       int
	Name     string
	Password string
	Role     string
}

type Users []*User

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (user *CreateUser) IsValid() error {
	var err error

	switch {
	case user.Name == "":
		err = apperror.ErrInvalidUserName
	case user.Password == "":
		err = apperror.ErrInvalidPassword
	case user.Role != "admin" && user.Role != "employee" && user.Role != "user":
		err = apperror.ErrInvalidRole
	}

	return err
}
