package loader

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"mycleanarch/internal/utils/ilog"
)

func InitDB(logger ilog.Logger) (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DATABASE_DRIVER"), dbURL())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if os.Getenv("DATABASE_DEBUG") == "true" {
		boil.DebugMode = true

		boil.DebugWriter = logger.WriterLevel(logrus.DebugLevel)
	}

	return db, nil
}

func dbURL() string {
	if os.Getenv("SSL_DISABLE") == "true" {
		return os.Getenv("DATABASE_URL") + "?sslmode=disable"
	}

	return os.Getenv("DATABASE_URL")
}
