package config

import (
	"github.com/juniorleaoo/learn-go/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func createDatabaseFileIfNotExist(dbPath string) error {
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		file.Close()
	}

	return nil
}

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	err := createDatabaseFileIfNotExist(dbPath)
	if err != nil {
		logger.Errorf("Error on creating database file %v", err)
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
		return nil, err
	}

	return db, nil
}
