package luyous

import (
	"changliang/zf"
	"changliang/zfzhi"
	"github.com/astaxie/beego"
	"mhsydata/zdkongzhiqis"
)

func init() {
	beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Juese(true), &zdkongzhiqis.Juesekongzhiqi{})
	beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Juese(true)+zfzhi.Zhi.Xx()+zfzhi.Zhi.Mh()+zf.Zfs.Id(false), &zdkongzhiqis.Juesekongzhiqi{})
	beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Juese(true)+zfzhi.Zhi.Xx()+zf.Zfs.Quanbu(true), &zdkongzhiqis.Jueseliebiaokongzhiqi{})

}
