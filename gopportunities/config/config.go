package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error on initialize database: %v", err)
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
