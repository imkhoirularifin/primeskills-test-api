package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func dbSetup() {
	var err error
	logMode := logger.Default.LogMode(logger.Silent)

	switch cfg.Database.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.New(sqlite.Config{
			DSN: cfg.Database.Dsn,
		}), &gorm.Config{
			Logger: logMode,
		})
	case "mysql":
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: cfg.Database.Dsn,
		}), &gorm.Config{
			Logger: logMode,
		})
	default:
		panic("Database driver not supported")
	}

	if err != nil {
		panic(err)
	}
}
