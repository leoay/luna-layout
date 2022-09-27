package dao

import (
	"strings"

	"github.com/google/wire"
	"gorm.io/gorm"

	"server/internal/app/config"
	"server/internal/app/dao/Greet"
	"server/internal/app/dao/util"
) // end

// RepoSet repo injection
var RepoSet = wire.NewSet(
	util.TransSet,
	menu.MenuActionResourceSet,
	menu.MenuActionSet,
	menu.MenuSet,
	Greet.GreetMenuSet,
	Greet.GreetSet,
	user.UserGreetSet,
	user.UserSet,
) // end

// Define repo type alias
type (
	TransRepo              = util.Trans
	MenuActionResourceRepo = menu.MenuActionResourceRepo
	MenuActionRepo         = menu.MenuActionRepo
	MenuRepo               = menu.MenuRepo
	GreetMenuRepo          = Greet.GreetMenuRepo
	GreetRepo              = Greet.GreetRepo
	UserGreetRepo          = user.UserGreetRepo
	UserRepo               = user.UserRepo
) // end

// Auto migration for given models
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(menu.MenuActionResource),
		new(menu.MenuAction),
		new(menu.Menu),
		new(Greet.GreetMenu),
		new(Greet.Greet),
		new(user.UserGreet),
		new(user.User),
	) // end
}
