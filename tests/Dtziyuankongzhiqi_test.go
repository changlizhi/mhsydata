package tests

import (
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"log"
	"mhsydata/moxings"
	"mhsydata/zdkongzhiqis"
	"strconv"
	"testing"
)

func TestPostdtziyuan(t *testing.T) {
	c := zdkongzhiqis.Dtziyuankongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	reqjson := zfzhi.Zhi.Postdtziyuan()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Post()
	log.Println(c.Data)

}
func TestPatchdtziyuan(t *testing.T) {
	c := zdkongzhiqis.Dtziyuankongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	reqjson := zfzhi.Zhi.Patchdtziyuan()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Patch()
	log.Println(c.Data)

}
func TestDeletedtziyuan(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	c := zdkongzhiqis.Dtziyuankongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Delete()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
func TestGetdtziyuan(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	c := zdkongzhiqis.Dtziyuankongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false), paramid)
	c.Get()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
func TestPostquanbu(t *testing.T) {
	c := zdkongzhiqis.Dtziyuanliebiaokongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	c.Post()
	zhis := c.Data[zf.Zfs.Json(true)]
	if zhi, ok := zhis.([]*moxings.Dtziyuan); ok {
		for _, z := range zhi {
			log.Println("z:====", z)
		}

	}

}
