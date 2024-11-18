package migrations

import (
	"ReactiveX22/code-nest-api/data"
	"context"
	"log"

	"github.com/uptrace/bun"
)

type Migration interface {
	Up(db *bun.DB) error
	Down(db *bun.DB) error
}

type UserMigration struct{}

func (m *UserMigration) Up(db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*data.User)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to run up migration: %v", err)
		return err
	}
	log.Println("User table created successfully")
	return nil
}

func (m *UserMigration) Down(db *bun.DB) error {
	_, err := db.NewDropTable().Model((*data.Session)(nil)).IfExists().Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to drop dependent table posts: %v", err)
		return err
	}
	log.Println("Dependent table sessions dropped successfully")

	_, err = db.NewDropTable().Model((*data.Post)(nil)).IfExists().Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to drop dependent table posts: %v", err)
		return err
	}
	log.Println("Dependent table posts dropped successfully")

	_, err = db.NewDropTable().Model((*data.User)(nil)).IfExists().Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to run down migration for user table: %v", err)
		return err
	}
	log.Println("User table dropped successfully")

	return nil
}
