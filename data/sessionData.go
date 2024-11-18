package data

import (
	"ReactiveX22/code-nest-api/db"
	"ReactiveX22/code-nest-api/utils"
	"context"
	"database/sql"
	"errors"
	"log"
)

var ErrInvalidCredential = errors.New("invalid login credential")

func CreateSession(ctx context.Context, loginReq LoginRequest) (*Session, error) {
	user := &User{}
	err := db.DB.NewSelect().Model(user).Column("*").Where("email= ?", loginReq.Email).Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidCredential
		}
		return nil, err
	}

	isCorrect := utils.CheckPassword(user.PasswordHash, loginReq.Password)

	if !isCorrect {
		return nil, ErrInvalidCredential
	}

	_, err = db.DB.NewDelete().Model(&Session{}).Where("user_id= ?", user.ID).Exec(ctx)
	if err != nil {
		log.Printf("Error deleting old session: %v", err)
		return nil, err
	}

	token, expireAt, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return nil, err
	}

	session := &Session{
		Token:    token,
		UserID:   user.ID,
		ExpireAt: expireAt,
	}

	_, err = db.DB.NewInsert().Model(session).Exec(ctx)

	if err != nil {
		log.Printf("Error login.")
		return nil, err
	}

	return session, nil
}

func DeleteSession(ctx context.Context, sessionToken string) error {
	result, err := db.DB.NewDelete().
		Model(&Session{}).
		Where("token = ?", sessionToken).
		Exec(ctx)
	if err != nil {
		log.Printf("Error deleting session with token %s: %v", sessionToken, err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No session found for token %s", sessionToken)
		return nil
	}

	return nil
}
