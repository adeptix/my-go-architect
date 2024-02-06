package main

import (
	"context"
	"database/sql"
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/bonch.dev/go-lib/migrator/grammar/postgres"
	"gitlab.com/bonch.dev/go-lib/migrator/migrator"

	"mycleanarch/internal/loader"
	_ "mycleanarch/internal/migrations"
	"mycleanarch/internal/utils/ilog"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Warning(err)
	}

	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := loader.InitLogger()

	db, err := loader.InitDB(logger.WithField("Type", "Database"))
	if err != nil {
		logger.WithError(err).WithField("Stage", "Migration").Fatal("DB initialization failed")
	}

	m := initMigrator(db, logger)

	freshMigrateFlag := flag.Bool("fresh", false, "Fresh and migrate")
	flag.Parse()

	if *freshMigrateFlag {
		freshDB(logger.WithField("Type", "DB Fresh"), m)
	}

	migrateDB(appCtx, logger.WithField("Type", "DB Migration"), m)
}

func migrateDB(ctx context.Context, logger ilog.Logger, m *migrator.Migrator) {
	logger.Info("DB Migration started")
	defer logger.Info("DB Migration ended")

	if err := m.MigrateCtx(ctx); err != nil {
		logger.WithError(err).Fatal("Migration")
	}
}

func freshDB(logger ilog.Logger, m *migrator.Migrator) {
	logger.Info("DB Fresh started")
	defer logger.Info("DB Fresh ended")

	if err := m.Wipe(); err != nil {
		logger.Panic(err.Error())
	}
}

func initMigrator(db *sql.DB, logger ilog.Logger) *migrator.Migrator {
	g := new(postgres.SQLGrammar)
	g.AddModifiers(postgres.DefaultModifiers()...)

	m := migrator.Init(db, g, logger)

	// Disabling persistent data migrations, if test running.
	if os.Getenv("APP_ENV") == "test" {
		m.DisablePersistent()
	}

	return m
}
