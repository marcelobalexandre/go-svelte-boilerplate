package user

import (
	"backend/internal/modules"
	"context"
	"database/sql"
	"os"
	"strings"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateUser func(context.Context, AuthenticateUserInput) (*AuthenticateUserOutput, error)

type AuthenticateUserInput struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type AuthenticateUserOutput struct {
	User  *User  `json:"-"`
	Token string `json:"token"`
}

func NewAuthenticateUser(
	userRepo UserRepo,
	validate *validator.Validate,
	trans ut.Translator,
) AuthenticateUser {
	return func(ctx context.Context, input AuthenticateUserInput) (*AuthenticateUserOutput, error) {
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

		user := User{}
		err = userRepo.db.QueryRow(ctx, "SELECT id, username, password_hash, created_at, updated_at FROM users WHERE username = ? AND deleted_at IS NULL", input.Username).
			Scan(&user.Id, &user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
		if err == sql.ErrNoRows {
			return nil, modules.ValidationError{
				Message: "Validation error",
			}
		} else if err != nil {
			return nil, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
		if err != nil {
			return nil, modules.ValidationError{
				Message: "Validation error",
			}
		}

		token, err := generateToken(user.Id)
		if err != nil {
			return nil, err
		}

		output := AuthenticateUserOutput{
			User:  &user,
			Token: token,
		}

		return &output, nil
	}
}

type TokenClaims struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

func generateToken(userId int64) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour)
	claims := &TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET")

	return token.SignedString([]byte(secret))
}
