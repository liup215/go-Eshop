package home

import (
	"github.com/tango-contrib/renders"
	"fmt"
	"marketing/common"
)
type IndexAction struct {
	renders.Renderer
}


//获取手机端小导航
var sn1 = map[string]string{
	"href":"#",
	"img_src":"images/navsmall.jpg",
	"title":"商品分类",
}

var sn2 = map[string]string{
	"href":"#",
	"img_src":"images/huismall.jpg",
	"title":"活动",
}

var sn3 = map[string]string{
	"href":"#",
	"img_src":"images/mansmall.jpg",
	"title":"个人中心",
}

var sn4 = map[string]string{
	"href":"#",
	"img_src":"images/moneysmall.jpg",
	"title":"我的钱包",
}

var smallnav = [...]map[string]string{sn1,sn2,sn3,sn4}

//获取商城头条
var mq1 = map[string]interface{}{
	"href":"#",
	"type":"特惠",
	"title":"商城爆品1分秒",
	"img_src":"images/TJ.jpg",
	"short":"xxxxxxxxxx",
}
var mq2 = map[string]interface{}{
	"href":"#",
	"type":"公告",
	"title":"商城与广州市签署战略合作协议",
	"img_src":"images/TJ2.jpg",
	"short":"**************",
}
var mq3 = map[string]interface{}{
	"href":"#",
	"type":"特惠",
	"title":"洋河年末大促，低至两件五折",
	"img_src":"images/TJ.jpg",
	"short":"xxxxxxxxxx",
}
var mq4 = map[string]interface{}{
	"href":"#",
	"type":"公告",
	"title":"华北、华中部分地区配送延迟",
	"img_src":"images/TJ2.jpg",
	"short":"**************",
}
var mq5 = map[string]interface{}{
	"href":"#",
	"type":"特惠",
	"title":"家电狂欢千亿礼券 买1送1！",
	"img_src":"images/TJ.jpg",
	"short":"xxxxxxxxxx",
}
var marqueen = [...]map[string]interface{}{mq1,mq2,mq3,mq4,mq5}

//获取商城主页内容
//1.今日推荐
var dayRec = map[string]interface{}{
	"topimg_src":"images/2016.png",
	"mTitle":"今日",
	"sTitle":"推荐",
	"content":[...]map[string]string{
		map[string]string{
			"mTitle":"真的有鱼",
			"sTitle":"开年福利篇",
			"href":"#",
			"img_src":"images/tj.png",
		},
		map[string]string{
			"mTitle":"囤货过冬",
			"sTitle":"让爱早回家",
			"href":"#",
			"img_src":"images/tj1.png",
		},
		map[string]string{
			"mTitle":"浪漫情人节",
			"sTitle":"甜甜蜜蜜",
			"href":"#",
			"img_src":"images/tj2.png",
		},

	},
}

//2.营销活动
var activity = map[string]interface{}{
	"mTitle":"活动",
	"sTitle":"每期活动 优惠享不停",
	"allAct":"#",
	"allalign":"全部活动",
	"content":[...]map[string]string{
		map[string]string{
			"style":"icon-sale one",
			"type":"秒杀",
			"img_src":"images/activity1.jpg",
			"title":"春节送礼优选",
			"href":"#",
		},
		map[string]string{
			"style":"icon-sale two",
			"type":"特惠",
			"img_src":"images/activity2.jpg",
			"title":"春节送礼优选",
			"href":"#",
		},
		map[string]string{
			"style":"icon-sale three",
			"type":"团购",
			"img_src":"images/activity3.jpg",
			"title":"春节送礼优选",
			"href":"#",
		},
		map[string]string{
			"style":"icon-sale",
			"type":"超值",
			"img_src":"images/activity.jpg",
			"title":"春节送礼优选",
			"href":"#",
		},
	},

}

//3.首页楼层
var floors = [...]map[string]interface{}{
	map[string]interface{}{
		"floor":"1",
		"mTitle":"甜品",
		"sTitle":"每一道甜品都有一个故事",
		"brands":[...]map[string]string{
			map[string]string{
				"href":"#",
				"name":"桂花糕",
			},
			map[string]string{
				"href":"#",
				"name":"奶皮酥",
			},
			map[string]string{
				"href":"#",
				"name":"栗子糕",
			},
			map[string]string{
				"href":"#",
				"name":"马卡龙",
			},
			map[string]string{
				"href":"#",
				"name":"铜锣烧",
			},
			map[string]string{
				"href":"#",
				"name":"豌豆黄",
			},
		},
		"adds":map[string]interface{}{
			"words":[...]map[string]string{
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
			},
			"href":"#",
			"img_src":"images/act1.png",
		},
		"floors":[...]map[string]string{
			map[string]string{
				"style":"am-u-sm-7 am-u-md-4 text-two sug",
				"title":"雪之恋和风大福",
				"price":"13.8",
				"href":"#",
				"img_src":"images/2.jpg",
			},
			map[string]string{
				"style":"am-u-sm-7 am-u-md-4 text-two",
				"title":"雪之恋和风大福",
				"price":"13.8",
				"href":"#",
				"img_src":"images/1.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three big",
				"title":"小优布丁",
				"price":"4.8",
				"href":"#",
				"img_src":"images/5.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three sug",
				"title":"小优布丁",
				"price":"4.8",
				"href":"#",
				"img_src":"images/3.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three",
				"title":"小优布丁",
				"price":"4.8",
				"href":"#",
				"img_src":"images/4.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three last big",
				"title":"小优布丁",
				"price":"4.8",
				"href":"#",
				"img_src":"images/5.jpg",
			},
		},


	},
	map[string]interface{}{
		"floor":"2",
		"mTitle":"甜品",
		"sTitle":"每一道甜品都有一个故事",
		"brands":[...]map[string]string{
			map[string]string{
				"href":"#",
				"name":"桂花糕",
			},
			map[string]string{
				"href":"#",
				"name":"奶皮酥",
			},
			map[string]string{
				"href":"#",
				"name":"栗子糕",
			},
			map[string]string{
				"href":"#",
				"name":"马卡龙",
			},
			map[string]string{
				"href":"#",
				"name":"铜锣烧",
			},
			map[string]string{
				"href":"#",
				"name":"豌豆黄",
			},
		},
		"adds":map[string]interface{}{
			"words":[...]map[string]string{
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
				map[string]string{
					"href":"#",
					"text":"核桃",
				},
			},
			"href":"#",
			"img_src":"images/act1.png",
		},
		"floors":[...]map[string]string{
			map[string]string{
				"style":"am-u-sm-7 am-u-md-4 text-two sug",
				"title":"雪之恋和风大福",
				"price":"1",
				"href":"#",
				"img_src":"images/2.jpg",
			},
			map[string]string{
				"style":"am-u-sm-7 am-u-md-4 text-two",
				"title":"雪之恋和风大福",
				"price":"2",
				"href":"#",
				"img_src":"images/1.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three big",
				"title":"小优布丁",
				"price":"3",
				"href":"#",
				"img_src":"images/5.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three sug",
				"title":"小优布丁",
				"price":"4",
				"href":"#",
				"img_src":"images/3.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three",
				"title":"小优布丁",
				"price":"5",
				"href":"#",
				"img_src":"images/4.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three last big",
				"title":"小优布丁",
				"price":"6",
				"href":"#",
				"img_src":"images/5.jpg",
			},
		},


	},
}


func (index *IndexAction) Get(){
	banners,_ := getBanners()
	topnav,_ := getTopNav()
	categories,_ := getCategories()
	var t = renders.T{
		"title":"全域营销商城",
		"banners":banners,
		"topnav":topnav,
		"categories":categories,
		"smallnav":smallnav,
		"marqueen":marqueen,
		"dayRec":dayRec,
		"activity":activity,
		"floors":floors,
	}
	index.Render("home/index.html",t)
}

func (index *IndexAction) Post(){

}

//获取banner
func getBanners()(banners []map[string]string,err error){

	banners,err = common.Query("select * from adds_list where type=1")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}

	return banners,nil

}

//获取首页导航
func getTopNav()(topnav []map[string]string,err error){
	topnav,err =common.Query("select * from top_nav")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}

	return topnav,nil
}

//获取分类导航
func getCategories()(categories []map[string]interface{},err error){
	topCat,err := common.Query("select id,name,align,img_src,parent_id from product_category where parent_id=0 and state=1 and nav=1 order by pc_sort")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}

	for _,value:= range topCat{
		var cat=make(map[string]interface{})

		//获取一级分类内容
		cat["cateTitle"]=map[string]interface{}{

			"img_src":value["img_src"],
			"title":value["name"],
			"align":value["align"],
		}

		fmt.Println("-------",value["id"])
		//获取二级分类内容
		var cat_2 []map[string]interface{}
		var cat_2_s []map[string]string
		cat_2_s,err := common.Query("select id,name,align,parent_id from product_category where parent_id="+string(value["id"])+" and state=1 and nav=1 order by pc_sort")
		if err!=nil{
			fmt.Println(err)
		}

		//获取三级分类内容
		var cat_3 []map[string]string
		for _,v := range cat_2_s{
			stmt :="select align from product_category where parent_id="+(v["id"])+ " and state=1 and nav=1 order by pc_sort"
			cat_3,err = common.Query(stmt)
			cat_2=append(cat_2,map[string]interface{}{
				"align":v["align"],
				"content":cat_3,
			})
		}

		cat["sort"]=cat_2

		//获取实力商家内容
		var brands = make(map[string]interface{})
		brands["title"]="实力商家"
		brands_s,err := common.Query("select name,align,link from product_brands where cat="+value["id"]+" order by sort",)
		if err!=nil{
			fmt.Println(err)
		}
		brands["content"]=brands_s

		cat["brand"]=brands
		for i:=0;i<10;i++{
			categories=append(categories,cat)
		}

	}

	return categories,nil
}
