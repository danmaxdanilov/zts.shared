package sqlite

import "gorm.io/gorm"

type Config struct {
	DbPath     string      `mapstructure:"dbPath"` //for inmemory use "file::memory:?cache=shared"
	GormConfig gorm.Config `mapstructure:"gormConfig"`
}
