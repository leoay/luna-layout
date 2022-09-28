package dao

import (
	"strings"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/leoay/luna/internal/app/dao/util"
	"luna-layout/internal/app/config"
	"luna-layout/internal/app/dao/Greet"
) // end

// RepoSet repo injection
var RepoSet = wire.NewSet(
	util.TransSet,
	Greet.GreetSet,
) // end

// Define repo type alias
type (
	TransRepo = util.Trans
	GreetRepo = Greet.GreetRepo
) // end

// AutoMigrate Auto migration for given models
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(Greet.Greet),
	) // end
}
