package app

import (
	"errors"
	"gorm.io/gorm"

	"luna-layout/internal/app/config"
	"luna-layout/internal/app/dao"
	"luna-layout/pkg/gormx"
)

func InitGormDB() (*gorm.DB, func(), error) {
	cfg := config.C.Gorm
	db, err := NewGormDB()
	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {}

	if cfg.EnableAutoMigrate {
		err = dao.AutoMigrate(db)
		if err != nil {
			return nil, cleanFunc, err
		}
	}

	return db, cleanFunc, nil
}

func NewGormDB() (*gorm.DB, error) {
	cfg := config.C
	var dsn string
	switch cfg.Gorm.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	default:
		return nil, errors.New("unknown db")
	}

	return gormx.New(&gormx.Config{
		Debug:        cfg.Gorm.Debug,
		DBType:       cfg.Gorm.DBType,
		DSN:          dsn,
		MaxIdleConns: cfg.Gorm.MaxIdleConns,
		MaxLifetime:  cfg.Gorm.MaxLifetime,
		MaxOpenConns: cfg.Gorm.MaxOpenConns,
		TablePrefix:  cfg.Gorm.TablePrefix,
	})
}
