package seeds

import (
	"database/sql"
	"log"
	"log/slog"

	"github.com/joho/godotenv"

	"github.com/webaiz/exercise7/blogging-platform/internal/db"
)

type Seed interface {
	Populate() error
}

type seeder struct {
	cleanups []func() error
	db       *sql.DB
}

func New() (Seed, error) {
	_ = godotenv.Load()

	d, err := db.NewPgSQL()

	if err != nil {
		slog.Error("database connection failed", "error", err)
		return nil, err
	}

	return &seeder{
		db:       d,
		cleanups: []func() error{},
	}, nil
}

func (s *seeder) Populate() error {
	tx, err := s.db.Begin()
	if err != nil {
		slog.Error("failed to start transaction", "error", err)
		return err
	}

	if err := s.startSeeding(tx); err != nil {
		slog.Error("failed to start seeding", "error", err)
		err := tx.Rollback()
		slog.Error("failed to rollback", "error", err)
		if err != nil {
			log.Fatal(err)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("failed to commit", "error", err)
		err := tx.Rollback()
		slog.Error("failed to rollback", "error", err)
		if err != nil {
			log.Fatal(err)
		}
		return err
	}

	for _, cleanup := range s.cleanups {
		if err := cleanup(); err != nil {
			slog.Error("failed to cleanup", "error", err)
			return err
		}
	}

	if err := s.db.Close(); err != nil {
		slog.Error("failed seeding closing db", slog.String("error", err.Error()))
		return err
	}

	slog.Info("seeding complete!")

	return nil
}

func (s *seeder) startSeeding(tx *sql.Tx) error {
	if err := s.post(tx); err != nil {
		slog.Error("failed seeding post", slog.String("error", err.Error()))
		return err
	}

	return nil
}
