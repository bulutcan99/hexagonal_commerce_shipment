package util

import (
	"github.com/bulutcan99/shipment/internal/core/domain"
	"github.com/oklog/ulid/v2"
	"reflect"
	"regexp"
)

func EmailValidator(email string) *domain.Error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return &domain.Error{
			Code:    domain.ErrInvalidEmail,
			Message: "Invalid email address",
		}
	}
	return nil
}

func RoleValidator(role string) *domain.Error {
	if role != "admin" && role != "user" {
		return &domain.Error{
			Code:    domain.ErrInvalidRole,
			Message: "Invalid role",
		}
	}
	return nil
}

func PasswordValidator(password string) *domain.Error {
	if len(password) < 9 {
		return &domain.Error{
			Code:    domain.ErrInvalidPassword,
			Message: "Password must be at least 9 characters long",
		}
	}
	if ok, _ := regexp.MatchString(`[A-Z]`, password); !ok {
		return &domain.Error{
			Code:    domain.ErrInvalidPassword,
			Message: "Password not meet requirements!",
		}
	}
	if ok, _ := regexp.MatchString(`[a-z]`, password); !ok {
		return &domain.Error{
			Code:    domain.ErrInvalidPassword,
			Message: "Password not meet requirements!",
		}
	}
	if ok, _ := regexp.MatchString(`[0-9]`, password); !ok {
		return &domain.Error{
			Code:    domain.ErrInvalidPassword,
			Message: "Password not meet requirements!",
		}
	}
	return nil
}

func TokenValidator(v interface{}, param string) *domain.Error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return &domain.Error{
			Code:    domain.InvalidToken,
			Message: "Token validation error!",
		}
	}
	_, err := ulid.Parse(v.(string))
	if err != nil {
		return &domain.Error{
			Code:    domain.InvalidToken,
			Message: "Token validation error!",
		}
	}
	return nil
}
