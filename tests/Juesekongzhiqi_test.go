package tests

import (
	"mhsydata/zdkongzhiqis"
	"mhsydata/moxings"
	"gongju"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"strconv"
	"testing"
)

func TestPostjuese(t*testing.T) {
	c := zdkongzhiqis.Juesekongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	reqjson := zfzhi.Zhi.Postjuese()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Post()
	log.Println(c.Data)

}
func TestPatchjuese(t*testing.T) {
	c := zdkongzhiqis.Juesekongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	reqjson := zfzhi.Zhi.Patchjuese()
	c.Ctx.Input.RequestBody = []byte(reqjson)
	c.Patch()
	log.Println(c.Data)

}
func TestDeletejuese(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	c := zdkongzhiqis.Juesekongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh() + zf.Zfs.Id(false), paramid)
	c.Delete()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
func TestGetjuese(t *testing.T) {
	paramid := strconv.Itoa(zfzhi.Zhi.Shuzi1())
	c := zdkongzhiqis.Juesekongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	c.Ctx.Input.SetParam(zfzhi.Zhi.Mh() + zf.Zfs.Id(false), paramid)
	c.Get()
	log.Println(c.Data[zf.Zfs.Json(true)])

}
func TestPostquanbu(t *testing.T) {
	c := zdkongzhiqis.Jueseliebiaokongzhiqi{}
	gongju.Kongzhiqi(&c.Controller)
	c.Post()
	zhis := c.Data[zf.Zfs.Json(true)]
	if zhi, ok := zhis.([]*moxings.Juese); ok {
		for _, z := range zhi {
			log.Println("z:====", z)
		}

	}

}
