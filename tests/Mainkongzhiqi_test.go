package tests

import (
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"log"
	"mhsydata/zdkongzhiqis"
	"strconv"
	"testing"
)

func TestGetmain(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	c := zdkongzhiqis.Mainkongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Get()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
