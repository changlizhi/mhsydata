package zdkongzhiqis

import (
	"changliang/zf"
	"changliang/zfzhi"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"mhsydata/chushihuas"
	"mhsydata/moxings"
	"mhsydata/zddtziyuanyewus"
	"strconv"
	"strings"
)

type Dtziyuankongzhiqi struct {
	beego.Controller
}
type Dtziyuanliebiaokongzhiqi struct {
	beego.Controller
}

func (c *Dtziyuankongzhiqi) Get() {
	canshu := c.GetString(zfzhi.Zhi.Mh() + zf.Zfs.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Zfs.Json(true)] = map[string]string{
			zf.Zfs.Error05(false): chushihuas.Cuowus[zf.Zfs.Error05(false)].Zhi,
		}
		c.ServeJSON()
		return
	}
	dtziyuan := zddtziyuanyewus.Chaxundtziyuan(id)
	c.Data[zf.Zfs.Json(true)] = dtziyuan
	c.ServeJSON()
	return
}
func (c *Dtziyuankongzhiqi) Post() {
	dtziyuan := moxings.Dtziyuan{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &dtziyuan)
	serviceret := zddtziyuanyewus.Tianjiadtziyuan(&dtziyuan)
	tishi := chushihuas.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = chushihuas.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Dtziyuankongzhiqi) Patch() {
	dtziyuan := moxings.Dtziyuan{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &dtziyuan)
	serviceret := zddtziyuanyewus.Xiugaidtziyuan(&dtziyuan)
	tishi := chushihuas.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = chushihuas.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Dtziyuankongzhiqi) Delete() {
	canshu := c.GetString(zfzhi.Zhi.Mh() + zf.Zfs.Id(false))
	id, err := strconv.Atoi(canshu)
	if err != nil {
		log.Println(err)
		c.Data[zf.Zfs.Json(true)] = map[string]string{
			zf.Zfs.Error05(false): chushihuas.Cuowus[zf.Zfs.Error05(false)].Zhi,
		}
		c.ServeJSON()
		return
	}
	serviceret := zddtziyuanyewus.Shanchudtziyuan(id)
	tishi := chushihuas.Tishis[serviceret].Zhi
	if tishi == zfzhi.Zhi.Kzf() {
		splitret := strings.Split(serviceret, zfzhi.Zhi.Xhx())
		c.Data[zf.Zfs.Json(true)] = chushihuas.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi + zfzhi.Zhi.Mh() + splitret[zfzhi.Zhi.Shuzi1()]
		c.ServeJSON()
		return
	}
	c.Data[zf.Zfs.Json(true)] = tishi
	c.ServeJSON()
	return
}
func (c *Dtziyuanliebiaokongzhiqi) Post() {
	ret := zddtziyuanyewus.Chaxunquanbudtziyuan()
	allret := make(map[string][]*moxings.Dtziyuan)
	fubianmas := make(map[string]string)
	for _, r := range ret {
		fubianmas[r.Fubianma] = r.Fubianma
	}
	for _, fubianma := range fubianmas {
		bm := []*moxings.Dtziyuan{}
		for _, r := range ret {
			if r.Fubianma == fubianma {
				bm = append(bm, r)
			}
		}
		allret[fubianma] = bm
	}
	c.Data[zf.Zfs.Json(true)] = allret
	c.ServeJSON()
	return
}
