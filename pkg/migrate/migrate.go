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

	migrationManager := migrations.NewMigrationManager()

	if *upFlag {
		err := migrationManager.RunMigrations(db, true)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migrations completed successfully.")
	} else if *downFlag {
		err := migrationManager.RunMigrations(db, false)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migrations rolled back successfully.")
	} else {
		log.Println("No migration was run. Use the --up or --down flag to run migrations.")
	}
}
