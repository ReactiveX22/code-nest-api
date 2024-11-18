package migrations

import (
	"ReactiveX22/code-nest-api/data"
	"context"
	"log"

	"github.com/uptrace/bun"
)

type SessionMigration struct{}

func (m *SessionMigration) Up(db *bun.DB) error {
	_, err := db.NewCreateTable().Model((*data.Session)(nil)).IfNotExists().ForeignKey("(user_id) REFERENCES users(id) ON DELETE CASCADE").Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to run up migration: %v", err)
		return err
	}
	_, err = db.NewCreateIndex().Model((*data.Session)(nil)).Unique().Column("token", "user_id").Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to run up migration: %v", err)
		return err
	}
	log.Println("Sessions table created successfully")
	return nil
}

func (m *SessionMigration) Down(db *bun.DB) error {
	_, err := db.NewDropTable().Model((*data.Session)(nil)).IfExists().Exec(context.Background())
	if err != nil {
		log.Fatalf("Failed to run down migration: %v", err)
		return err
	}
	log.Println("Sessions table dropped successfully")
	return nil
}
