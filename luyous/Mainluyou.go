package luyous

import (
	"changliang/zfzhi"
	"changliang/zh"
	"github.com/astaxie/beego"
	"mhsydata/zdkongzhiqis"

	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter(
		zfzhi.Zhi.Xh(),
		beego.BeforeRouter, cors.Allow(
			&cors.Options{
				AllowAllOrigins:  true,
				AllowMethods:     zh.Zhs.AllowMethods(),
				AllowHeaders:     zh.Zhs.AllowHeaders(),
				ExposeHeaders:    zh.Zhs.ExposeHeaders(),
				AllowCredentials: true,
			},
		),
	)
	beego.Router(
		zfzhi.Zhi.Xx(),
		&zdkongzhiqis.Mainkongzhiqi{},
	)

}
