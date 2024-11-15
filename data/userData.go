package data

import (
	"ReactiveX22/code-nest-api/db"
	"ReactiveX22/code-nest-api/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

var ErrUserNotFound = errors.New("user not found")

func CreateUser(ctx context.Context, createUserReq CreateUserRequest) (*User, error) {
	hashedPassword, err := utils.HashPassword(createUserReq.Password)

	if err != nil {
		log.Printf("Error hashing password.")
		return nil, err
	}

	user := &User{
		Username:     createUserReq.Username,
		Email:        createUserReq.Email,
		PasswordHash: hashedPassword,
	}

	_, err = db.DB.NewInsert().Model(user).Exec(ctx)

	if err != nil {
		log.Printf("Error creating user.")
		return nil, err
	}

	return user, err
}

func GetUserByID(ctx context.Context, id int) (*User, error) {
	user := &User{}

	err := db.DB.NewSelect().Model(user).Where("id= ?", id).Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func UpdateUser(ctx context.Context, id int, updateUserReq UpdateUserRequest) (*User, error) {
	var passwordHash string

	if updateUserReq.Password != "" {
		var err error
		passwordHash, err = utils.HashPassword(updateUserReq.Password)
		if err != nil {
			return nil, err
		}
	}

	userToUpdate := &User{
		Username:     updateUserReq.Username,
		Email:        updateUserReq.Email,
		PasswordHash: passwordHash,
		UpdatedAt:    time.Now(),
	}

	updatedUser := new(User)
	err := db.DB.NewUpdate().
		Model(userToUpdate).
		ExcludeColumn("created_at").
		Where("id = ?", id).
		OmitZero().
		Returning("*").
		Scan(ctx, updatedUser)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func DeleteUser(ctx context.Context, id int) error {
	user := new(User)
	result, err := db.DB.NewDelete().Model(user).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}
	return nil
}
