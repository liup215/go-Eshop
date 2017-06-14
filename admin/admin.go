package admin

import (
	"github.com/tango-contrib/renders"
	"fmt"
	"github.com/tango-contrib/binding"
	"github.com/lunny/tango"
	//"github.com/eknkc/amber/parser"
)

type LoginAction struct {
	renders.Renderer
	binding.Binder
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

func (login *LoginAction) Post() interface{}{
	//type mydata struct {
	//	Username string `json:"Username"`
	//	Pass string `json:"Password"`
	//}
	//var data mydata
	login.Req().ParseForm()
	fmt.Println(login.Req().PostForm.Get("Username"))
	return "aaa"
}
