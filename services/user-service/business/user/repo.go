package user

import (
	"context"
	"database/sql"
	modelUser "user-service/models/user"

	"github.com/jackc/pgx/v4/pgxpool"
)

var _ UserStorer = (*UserRepo)(nil)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewRepoUser(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateOne(ctx context.Context, user modelUser.UserEntity) error {
	const query = `INSERT INTO users (user_id,phone_number,email,password,name,photo_profile) 
			VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := u.db.Exec(ctx, query,
		user.UserId,
		user.PhoneNumber,
		user.Email,
		user.Password,
		user.Name,
		user.PhotoProfile,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) GetByEmailOrPhoneNumber(ctx context.Context, username string) (modelUser.UserEntity, error) {
	var (
		query  = `SELECT user_id, password FROM users WHERE deleted_at is null AND (email = $1 OR phone_number = $2)`
		result modelUser.UserEntity
	)

	err := u.db.QueryRow(ctx, query, username, username).Scan(
		&result.UserId,
		&result.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, sql.ErrNoRows
		}
		return result, err
	}

	return result, nil
}
