package user

import (
	"backend/internal/modules"
	"context"
	"strings"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser func(context.Context, CreateUserInput) (*User, error)

type CreateUserInput struct {
	Username string `validate:"required,lowercase"`
	Password string `validate:"required,gte=6"`
}

func NewCreateUser(
	userRepo UserRepo,
	validate *validator.Validate,
	trans ut.Translator,
) CreateUser {
	return func(ctx context.Context, input CreateUserInput) (*User, error) {
		err := validate.Struct(input)
		if err != nil {
			details := make(map[string][]string)
			errs := err.(validator.ValidationErrors)
			for _, err := range errs {
				field := strings.ToLower(err.Field())
				details[field] = append(details[field], err.Translate(trans))
			}

			return nil, modules.ValidationError{
				Message: "Validation error",
				Details: details,
			}
		}

		now := time.Now().UTC()
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		user := &User{
			Username:     input.Username,
			PasswordHash: string(passwordHash),
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		user, err = userRepo.Store(ctx, user)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}
