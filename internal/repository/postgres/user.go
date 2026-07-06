package postgres

import (
	"context"
	"quran-memory-planner/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
	INSERT INTO users(
		email, 
		password_hash, 
		name, 
		role
	)
	VALUES (
		$1,
		$2,
		$3,
		$4
	)
	RETURNING id
	`

	err := r.db.QueryRow(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.Role,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}
