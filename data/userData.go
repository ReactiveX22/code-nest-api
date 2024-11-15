package data

import (
	"ReactiveX22/code-nest-api/db"
	"ReactiveX22/code-nest-api/utils"
	"context"
	"log"
)

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
