package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Migration represents a single database migration
type Migration struct {
	Name      string
	SQL       string
	IsApplied bool
}

// MigrationManager handles database migrations
type MigrationManager struct {
	db *sql.DB
}

// NewMigrationManager creates a new migration manager
func NewMigrationManager(db *sql.DB) (*MigrationManager, error) {
	// Create migrations table if it doesn't exist
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version TEXT PRIMARY KEY,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return nil, fmt.Errorf("error creating migrations table: %v", err)
	}

	return &MigrationManager{db: db}, nil
}

// LoadMigrations loads all migration files from the specified directory
func (m *MigrationManager) LoadMigrations(migrationsDir string) ([]Migration, error) {
	var migrations []Migration

	// Read all .sql files from migrations directory
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return nil, fmt.Errorf("error reading migrations directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			content, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("error reading migration file %s: %v", file.Name(), err)
			}

			migrations = append(migrations, Migration{
				Name: file.Name(),
				SQL:  string(content),
			})
		}
	}

	// Sort migrations by filename
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Name < migrations[j].Name
	})

	return migrations, nil
}

// CheckAppliedMigrations marks migrations that have already been applied
func (m *MigrationManager) CheckAppliedMigrations(migrations []Migration) error {
	for i := range migrations {
		var exists bool
		err := m.db.QueryRow("SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)",
			migrations[i].Name).Scan(&exists)
		if err != nil {
			return fmt.Errorf("error checking migration %s: %v", migrations[i].Name, err)
		}
		migrations[i].IsApplied = exists
	}
	return nil
}

// ApplyMigrations applies all pending migrations
func (m *MigrationManager) ApplyMigrations(migrations []Migration) error {
	for _, migration := range migrations {
		if !migration.IsApplied {
			log.Printf("Applying migration: %s", migration.Name)

			// Start transaction
			tx, err := m.db.Begin()
			if err != nil {
				return fmt.Errorf("error starting transaction for %s: %v", migration.Name, err)
			}

			// Apply migration
			if _, err := tx.Exec(migration.SQL); err != nil {
				tx.Rollback()
				return fmt.Errorf("error applying migration %s: %v", migration.Name, err)
			}

			// Record migration
			if _, err := tx.Exec("INSERT INTO schema_migrations (version) VALUES ($1)",
				migration.Name); err != nil {
				tx.Rollback()
				return fmt.Errorf("error recording migration %s: %v", migration.Name, err)
			}

			// Commit transaction
			if err := tx.Commit(); err != nil {
				return fmt.Errorf("error committing migration %s: %v", migration.Name, err)
			}

			log.Printf("Successfully applied migration: %s", migration.Name)
		}
	}
	return nil
}

// RunMigrations runs all pending migrations
func RunMigrations(db *sql.DB, migrationsDir string) error {
	manager, err := NewMigrationManager(db)
	if err != nil {
		return fmt.Errorf("error creating migration manager: %v", err)
	}

	migrations, err := manager.LoadMigrations(migrationsDir)
	if err != nil {
		return fmt.Errorf("error loading migrations: %v", err)
	}

	if err := manager.CheckAppliedMigrations(migrations); err != nil {
		return fmt.Errorf("error checking applied migrations: %v", err)
	}

	if err := manager.ApplyMigrations(migrations); err != nil {
		return fmt.Errorf("error applying migrations: %v", err)
	}

	return nil
}
