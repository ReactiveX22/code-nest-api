package main

import (
	"ReactiveX22/code-nest-api/db"
	"ReactiveX22/code-nest-api/pkg/migrate/migrations"
	"flag"
	"log"
)

func main() {
	upFlag := flag.Bool("up", false, "Run migrations")
	downFlag := flag.Bool("down", false, "Run migrations")
	flag.Parse()

	db.InitDB()
	db := db.DB
	userMigration := &migrations.UserMigration{}

	if *upFlag {
		err := userMigration.Up(db)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migration completed successfully.")
	} else if *downFlag {
		err := userMigration.Down(db)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migration completed successfully.")
	} else {
		log.Println("No migration was run. Use the --up or --down flag to run migrations.")
	}
}
