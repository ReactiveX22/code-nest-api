package migrations

import (
	"ReactiveX22/code-nest-api/data"
	"context"
	"log"

	"github.com/uptrace/bun"
)

type PostMigration struct{}

func (m *PostMigration) Up(db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*data.Post)(nil)).IfNotExists().ForeignKey("(author_id) REFERENCES users(id) ON DELETE CASCADE").Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to run up migration: %v", err)
		return err
	}
	log.Println("Post table created successfully")
	return nil
}

func (m *PostMigration) Down(db *bun.DB) error {
	_, err := db.NewDropTable().Model((*data.Post)(nil)).IfExists().Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to run down migration: %v", err)
		return err
	}
	log.Println("Post table dropped successfully")
	return nil
}
