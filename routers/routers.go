package routers

import (
	"github.com/lunny/tango"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/binding"
	"marketing/home"
	"marketing/admin"
)

func InitTango()(*tango.Tango){
	t:=tango.Classic();
	t.Use(tango.Static())
	t.Use(renders.New(renders.Options{
		Reload:true,
		Directory:"./templates",
	}))
	t.Use(binding.Bind())

	t.Get("/",new(home.IndexAction))
	t.Get("/admin/login",new(admin.LoginAction))
	t.Post("/admin/login",new(admin.LoginAction))

	return t
}

