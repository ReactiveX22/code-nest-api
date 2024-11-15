package data

import (
	"ReactiveX22/code-nest-api/db"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

var ErrPostNotFound = errors.New("post not found")

func GetPosts(ctx context.Context) ([]Post, error) {
	var posts []Post

	err := db.DB.NewSelect().
		Model(&posts).
		Order("created_at DESC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	for i := range posts {
		loadPostWithAuthor(ctx, &posts[i], int(posts[i].ID))
	}

	return posts, nil
}

func CreatePost(ctx context.Context, createPostReq CreatePostRequest) (*Post, error) {
	post := &Post{
		AuthorID: createPostReq.AuthorID,
		Title:    createPostReq.Title,
		Content:  createPostReq.Content,
	}

	err := db.DB.NewInsert().Model(post).Scan(ctx)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		return nil, err
	}

	post, err = loadPostWithAuthor(ctx, post, int(post.ID))
	if err != nil {
		return nil, err
	}

	return post, nil
}

func GetPostByID(ctx context.Context, id int) (*Post, error) {
	post := &Post{}

	err := db.DB.NewSelect().Model(post).Where("id= ?", id).Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	post, err = loadPostWithAuthor(ctx, post, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func UpdatePost(ctx context.Context, id int, updatePostReq UpdatePostRequest) (*Post, error) {
	postToUpdate := &Post{
		Title:     updatePostReq.Title,
		Content:   updatePostReq.Content,
		UpdatedAt: time.Now(),
	}

	updatedPost := new(Post)
	err := db.DB.NewUpdate().
		Model(postToUpdate).
		ExcludeColumn("created_at", "author_id").
		Where("id = ?", id).
		OmitZero().
		Returning("*").
		Scan(ctx, updatedPost)

	if err != nil {
		return nil, err
	}
	updatedPost, err = loadPostWithAuthor(ctx, updatedPost, id)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func DeletePost(ctx context.Context, id int) error {
	post := new(Post)
	result, err := db.DB.NewDelete().Model(post).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("post with id %d not found", id)
	}
	return nil
}

// helpers

func loadPostWithAuthor(ctx context.Context, post *Post, postID int) (*Post, error) {
	err := db.DB.NewSelect().
		Model(post).
		ModelTableExpr("posts AS p").
		Where("p.id = ?", postID).
		Relation("Author").
		Scan(ctx)
	if err != nil {
		log.Printf("Error loading post with author: %v", err)
		return nil, err
	}
	return post, nil
}
