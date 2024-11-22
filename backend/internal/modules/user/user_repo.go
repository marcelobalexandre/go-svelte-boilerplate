package user

import (
	"backend/internal/database"
	"context"
	"time"
)

type User struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserRepo struct {
	db database.Service
}

func NewUserRepo(db database.Service) UserRepo {
	return UserRepo{db}
}

func (repo UserRepo) Store(ctx context.Context, user *User) (*User, error) {
	err := repo.db.QueryRow(
		ctx,
		"INSERT INTO users (username, password_hash, created_at, updated_at) VALUES (?, ?, ?, ?) RETURNING id",
		user.Username,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
