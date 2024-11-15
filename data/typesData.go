package data

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID           int64     `bun:"id,pk,autoincrement" json:"id"`
	Username     string    `bun:"username,notnull,unique" json:"username"`
	Email        string    `bun:"email,notnull,unique" json:"email"`
	PasswordHash string    `bun:"password_hash,notnull" json:"-"`
	CreatedAt    time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt    time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updatedAt"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
