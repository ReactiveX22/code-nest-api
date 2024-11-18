package migrations

import "github.com/uptrace/bun"

type MigrationManager struct {
	migrations []Migration
}

func NewMigrationManager() *MigrationManager {
	return &MigrationManager{
		migrations: []Migration{
			&UserMigration{},
			&PostMigration{},
			&SessionMigration{},
		},
	}
}

func (m *MigrationManager) RunMigrations(db *bun.DB, up bool) error {
	for _, migration := range m.migrations {
		var err error
		if up {
			err = migration.Up(db)
		} else {
			err = migration.Down(db)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
