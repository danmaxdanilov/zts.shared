package sqlite

import (
	"github.com/danmaxdanilov/zts.shared/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ISqliteContext interface {
}

type sqliteContext struct {
	logger logger.Logger
	config Config
	db     gorm.DB
}

func NewDbContext(
	logger logger.Logger,
	config Config,
) *sqliteContext {
	return &sqliteContext{
		logger: logger,
		config: config,
		db:     *initDb(logger, config),
	}
}

func initDb(logger logger.Logger, config Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DbPath), &config.GormConfig)
	if err != nil {
		logger.Fatal("Couldn't initialize database")
		return nil
	}

	return db
}
