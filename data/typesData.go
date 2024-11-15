package data

import (
	"time"

	"github.com/uptrace/bun"
)

// User Types
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

// Post Types
type Post struct {
	bun.BaseModel `bun:"table:posts,alias:p"`

	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	AuthorID  int64     `bun:"author_id,notnull" json:"authorId"`
	Title     string    `bun:"title,notnull" json:"title"`
	Content   string    `bun:"content,notnull" json:"content"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	Author *User `bun:"rel:belongs-to,join:author_id=id,on_delete:cascade" json:"author"`
}

type CreatePostRequest struct {
	AuthorID int64  `json:"authorId"`
	Title    string `bun:"title,notnull" json:"title"`
	Content  string `json:"content"`
}
type UpdatePostRequest struct {
	AuthorID int64  `json:"authorId"`
	Title    string `bun:"title,notnull" json:"title"`
	Content  string `json:"content"`
}
