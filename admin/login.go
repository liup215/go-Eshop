package admin

import (
	"github.com/tango-contrib/renders"
	"github.com/lunny/tango"
)

type LoginAction struct {
	renders.Renderer
	tango.Json
	tango.Ctx
}

type Admin struct{
	Username string
	Password string
}
func (login *LoginAction) Get(){
	login.Render("admin/login.html",renders.T{
		"title":"全域营销",
	})
}

func (login *LoginAction) Post(){
	data:=new(Admin)
	login.Ctx.Req().ParseForm();
	data.Username=login.Ctx.Req().FormValue("username")
	data.Password=login.Ctx.Req().FormValue("password")

	login.Render("admin/index.html",renders.T{
		"title":"全域营销",
	})
}
