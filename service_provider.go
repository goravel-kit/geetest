package geetest

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "geetest"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewGeetest(app.MakeConfig()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/goravel-kit/geetest", map[string]string{
		"config/geetest.go": app.ConfigPath("geetest.go"),
	})
}
