package admin

import (
	"github.com/tango-contrib/renders"
	"marketing/common"
	"fmt"
)

type IndexAction struct{
	renders.Renderer
}

func(index *IndexAction) Get(){
	logo,_:=getLogo()
	t:=renders.T{
		"title":"圈里全域营销后台管理系统",
		"name":"全域营销",
		"logo":logo,
	}
	index.Render("admin/index.html",t)
}

func getLogo()(logo map[string]string,err error){
	logos,err := common.Query("select * from f_logo where type=1")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	logo= logos[0]
	return logo,nil

}