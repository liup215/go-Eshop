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
	"imgSrc":"images/navsmall.jpg",
	"title":"商品分类",
}

var sn2 = map[string]string{
	"href":"#",
	"imgSrc":"images/huismall.jpg",
	"title":"活动",
}

var sn3 = map[string]string{
	"href":"#",
	"imgSrc":"images/mansmall.jpg",
	"title":"个人中心",
}

var sn4 = map[string]string{
	"href":"#",
	"imgSrc":"images/moneysmall.jpg",
	"title":"我的钱包",
}

var smallnav = [...]map[string]string{sn1,sn2,sn3,sn4}

//获取商城头条
var mq1 = map[string]interface{}{
	"href":"#",
	"type":"特惠",
	"title":"商城爆品1分秒",
	"imgSrc":"images/TJ.jpg",
	"short":"xxxxxxxxxx",
}
var mq2 = map[string]interface{}{
	"href":"#",
	"type":"公告",
	"title":"商城与广州市签署战略合作协议",
	"imgSrc":"images/TJ2.jpg",
	"short":"**************",
}
var mq3 = map[string]interface{}{
	"href":"#",
	"type":"特惠",
	"title":"洋河年末大促，低至两件五折",
	"imgSrc":"images/TJ.jpg",
	"short":"xxxxxxxxxx",
}
var mq4 = map[string]interface{}{
	"href":"#",
	"type":"公告",
	"title":"华北、华中部分地区配送延迟",
	"imgSrc":"images/TJ2.jpg",
	"short":"**************",
}
var mq5 = map[string]interface{}{
	"href":"#",
	"type":"特惠",
	"title":"家电狂欢千亿礼券 买1送1！",
	"imgSrc":"images/TJ.jpg",
	"short":"xxxxxxxxxx",
}
var marqueen = [...]map[string]interface{}{mq1,mq2,mq3,mq4,mq5}

//获取商城主页内容
//1.今日推荐
var dayRec = map[string]interface{}{
	"topImgSrc":"images/2016.png",
	"mTitle":"今日",
	"sTitle":"推荐",
	"content":[...]map[string]string{
		map[string]string{
			"mTitle":"真的有鱼",
			"sTitle":"开年福利篇",
			"href":"#",
			"imgSrc":"images/tj.png",
		},
		map[string]string{
			"mTitle":"囤货过冬",
			"sTitle":"让爱早回家",
			"href":"#",
			"imgSrc":"images/tj1.png",
		},
		map[string]string{
			"mTitle":"浪漫情人节",
			"sTitle":"甜甜蜜蜜",
			"href":"#",
			"imgSrc":"images/tj2.png",
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
			"imgSrc":"images/activity1.jpg",
			"title":"春节送礼优选",
			"href":"#",
		},
		map[string]string{
			"style":"icon-sale two",
			"type":"特惠",
			"imgSrc":"images/activity2.jpg",
			"title":"春节送礼优选",
			"href":"#",
		},
		map[string]string{
			"style":"icon-sale three",
			"type":"团购",
			"imgSrc":"images/activity3.jpg",
			"title":"春节送礼优选",
			"href":"#",
		},
		map[string]string{
			"style":"icon-sale",
			"type":"超值",
			"imgSrc":"images/activity.jpg",
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
			"imgSrc":"images/act1.png",
		},
		"floors":[...]map[string]string{
			map[string]string{
				"style":"am-u-sm-7 am-u-md-4 text-two sug",
				"title":"雪之恋和风大福",
				"price":"13.8",
				"href":"#",
				"imgSrc":"images/2.jpg",
			},
			map[string]string{
				"style":"am-u-sm-7 am-u-md-4 text-two",
				"title":"雪之恋和风大福",
				"price":"13.8",
				"href":"#",
				"imgSrc":"images/1.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three big",
				"title":"小优布丁",
				"price":"4.8",
				"href":"#",
				"imgSrc":"images/5.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three sug",
				"title":"小优布丁",
				"price":"4.8",
				"href":"#",
				"imgSrc":"images/3.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three",
				"title":"小优布丁",
				"price":"4.8",
				"href":"#",
				"imgSrc":"images/4.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three last big",
				"title":"小优布丁",
				"price":"4.8",
				"href":"#",
				"imgSrc":"images/5.jpg",
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
			"imgSrc":"images/act1.png",
		},
		"floors":[...]map[string]string{
			map[string]string{
				"style":"am-u-sm-7 am-u-md-4 text-two sug",
				"title":"雪之恋和风大福",
				"price":"1",
				"href":"#",
				"imgSrc":"images/2.jpg",
			},
			map[string]string{
				"style":"am-u-sm-7 am-u-md-4 text-two",
				"title":"雪之恋和风大福",
				"price":"2",
				"href":"#",
				"imgSrc":"images/1.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three big",
				"title":"小优布丁",
				"price":"3",
				"href":"#",
				"imgSrc":"images/5.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three sug",
				"title":"小优布丁",
				"price":"4",
				"href":"#",
				"imgSrc":"images/3.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three",
				"title":"小优布丁",
				"price":"5",
				"href":"#",
				"imgSrc":"images/4.jpg",
			},
			map[string]string{
				"style":"am-u-sm-3 am-u-md-2 text-three last big",
				"title":"小优布丁",
				"price":"6",
				"href":"#",
				"imgSrc":"images/5.jpg",
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
	db,err := common.GetDB()
	rows,err:= db.Query("select * from adds_list where type=1")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}

	for rows.Next(){
		var id int
		var style string
		var bk_color string
		var link string
		var img_src string
		err = rows.Scan(&id,&style,&bk_color,&link,&img_src)
		if err!=nil{
			fmt.Println(err)
			return nil,err
		}
		banners = append(banners,map[string]string{
			"style":bk_color,
			"href":link,
			"imgSrc":img_src,
		})
	}
	fmt.Println(banners)
	return banners,nil

}

//获取首页导航
func getTopNav()(topnav []map[string]string,err error){
	db,err := common.GetDB()
	rows,err :=db.Query("select * from top_nav")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	for rows.Next(){
		var id int
		var name string
		var link string
		err = rows.Scan(&id,&name,&link)
		if err!=nil{
			fmt.Println(err)
			return nil,err
		}
		topnav = append(topnav,map[string]string{
			"name":name,
			"href":link,
		})
	}
	return topnav,nil
}

//获取分类导航
func getCategories()(categories []map[string]interface{},err error){
	db,err := common.GetDB()
	rows,err := db.Query("select id,name,align,img_src,parent_id from product_category where parent_id=0 and state=1 and nav=1 order by pc_sort")
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	var topCat []map[string]interface{}
	for rows.Next(){
		var id int
		var name string
		var align string
		var imgSrc string
		var parent_id int
		err = rows.Scan(&id,&name,&align,&imgSrc,&parent_id)
		if err!=nil{
			fmt.Println(err)
			return nil,err
		}
		topCat = append(topCat,map[string]interface{}{

			"id":id,
			"name":name,
			"align":align,
			"imgSrc":imgSrc,
			"parent_id":parent_id,
		})
	}

	fmt.Println(topCat)

	for _,value:= range topCat{
		var cat=make(map[string]interface{})

		//获取一级分类内容
		cat["cateTitle"]=map[string]interface{}{

			"imgSrc":value["imgSrc"],
			"title":value["name"],
			"align":value["align"],
		}

		fmt.Println("-------",value["id"])
		//获取二级分类内容
		var cat_2 []map[string]interface{}
		rows_2,err := db.Query("select id,name,align,parent_id from product_category where parent_id=? and state=1 and nav=1 order by pc_sort",value["id"])
		if err!=nil{
			fmt.Println(err)
		}

		for rows_2.Next(){
			var cat2id int
			var cat2name string
			var cat2align string
			var cat2parent_id int
			err = rows_2.Scan(&cat2id,&cat2name,&cat2align,&cat2parent_id)
			fmt.Println("-*_*_*_*_*_",cat2id)
			var cat_3 []string
			//获取三级分类内容
			row_3,err:=db.Query("select align from product_category where parent_id=? and state=1 and nav=1 order by pc_sort",cat2id)
			if err!=nil{
				fmt.Println(err)
			}
			for row_3.Next(){
				var cat3align string
				err = row_3.Scan(&cat3align)
				cat_3=append(cat_3,cat3align)
			}
			cat_2=append(cat_2,map[string]interface{}{
				"align":cat2align,
				"content":cat_3,
			})
		}
		fmt.Println("**************",cat_2)
		cat["sort"]=cat_2

		//获取实力商家内容
		var brands = make(map[string]interface{})
		brands["title"]="实力商家"
		var brandsContent []map[string]string
		rows,err = db.Query("select name,align,link from product_brands where cat=? order by sort",value["id"])
		if err!=nil{
			fmt.Println(err)
		}
		for rows.Next(){
			var brandsName string
			var brandsAlign string
			var brandsLink string
			err = rows.Scan(&brandsName,&brandsAlign,&brandsLink)
			for i:=0;i<10;i++{
				brandsContent=append(brandsContent,map[string]string{
					"title":brandsName,
					"name":brandsAlign,
					"href":brandsLink,
				})
			}

		}
		brands["content"]=brandsContent
		cat["brand"]=brands
		for i:=0;i<10;i++{
			categories=append(categories,cat)
		}

	}

	return categories,nil
}
