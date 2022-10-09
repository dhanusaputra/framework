package cache

import (
	"github.com/goravel/framework/cache/console"
	console2 "github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/facades"
)

type ServiceProvider struct {
}

func (database *ServiceProvider) Register() {
	app := Application{}
	facades.Cache = app.Init()
}

func (database *ServiceProvider) Boot() {
	database.registerCommands()
}

func (database *ServiceProvider) registerCommands() {
	facades.Artisan.Register([]console2.Command{
		&console.ClearCommand{},
	})
}
