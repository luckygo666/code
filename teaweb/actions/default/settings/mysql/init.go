package mysql

import (
	"github.com/TeaWeb/code/teaweb/actions/default/settings"
	"github.com/TeaWeb/code/teaweb/configs"
	"github.com/TeaWeb/code/teaweb/helpers"
	"github.com/iwind/TeaGo"
)

var shouldRestart = false

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(&helpers.UserMustAuth{
				Grant: configs.AdminGrantAll,
			}).
			Helper(new(settings.Helper)).
			Prefix("/settings/mysql").
			Get("", new(IndexAction)).
			GetPost("/update", new(UpdateAction)).
			GetPost("/test", new(TestAction)).
			Get("/data", new(DataAction)).
			Get("/clean", new(CleanAction)).
			GetPost("/cleanUpdate", new(CleanUpdateAction)).
			EndAll()
	})
}
