package zdkongzhiqis
import(
"encoding/json"
"github.com/astaxie/beego"
"mhsydata/moxings"
"mhsydata/chushihuas"
"mhsydata/zdjuesequanxianyewus"
"changliang/zf"
"changliang/zfzhi"
"log"
"strconv"
"strings"

)
type Juesequanxiankongzhiqi struct{
beego.Controller
}
func (c *Juesequanxiankongzhiqi)Get(){
canshu:=c.GetString(zfzhi.Zhi.Mh()+zf.Zfs.Id(false))
id,err:=strconv.Atoi(canshu)
if err!=nil{
log.Println(err)
c.Data[zf.Zfs.Json(true)]=map[string]string{
zf.Zfs.Error05(false):chushihuas.Cuowus[zf.Zfs.Error05(false)].Zhi,

}
c.ServeJSON()
return
}
juesequanxian:=zdjuesequanxianyewus.Chaxunjuesequanxian(id)
c.Data[zf.Zfs.Json(true)]=juesequanxian
c.ServeJSON()
return
}
func (c *Juesequanxiankongzhiqi)Post(){
juesequanxian:=moxings.Juesequanxian{}
json.Unmarshal(c.Ctx.Input.RequestBody,&juesequanxian)
serviceret:=zdjuesequanxianyewus.Tianjiajuesequanxian(&juesequanxian)
tishi:=chushihuas.Tishis[serviceret].Zhi
if tishi==zfzhi.Zhi.Kzf(){
splitret:=strings.Split(serviceret,zfzhi.Zhi.Xhx())
c.Data[zf.Zfs.Json(true)]=chushihuas.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi+zfzhi.Zhi.Mh()+splitret[zfzhi.Zhi.Shuzi1()]
c.ServeJSON()
return
}
c.Data[zf.Zfs.Json(true)]=tishi
c.ServeJSON()
return
}
func (c *Juesequanxiankongzhiqi)Patch(){
juesequanxian:=moxings.Juesequanxian{}
json.Unmarshal(c.Ctx.Input.RequestBody,&juesequanxian)
serviceret:=zdjuesequanxianyewus.Xiugaijuesequanxian(&juesequanxian)
tishi:=chushihuas.Tishis[serviceret].Zhi
if tishi==zfzhi.Zhi.Kzf(){
splitret:=strings.Split(serviceret,zfzhi.Zhi.Xhx())
c.Data[zf.Zfs.Json(true)]=chushihuas.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi+zfzhi.Zhi.Mh()+splitret[zfzhi.Zhi.Shuzi1()]
c.ServeJSON()
return
}
c.Data[zf.Zfs.Json(true)]=tishi
c.ServeJSON()
return
}
func (c *Juesequanxiankongzhiqi)Delete(){
canshu:=c.GetString(zfzhi.Zhi.Mh()+zf.Zfs.Id(false))
id,err:=strconv.Atoi(canshu)
if err!=nil{
log.Println(err)
c.Data[zf.Zfs.Json(true)]=map[string]string{
zf.Zfs.Error05(false):chushihuas.Cuowus[zf.Zfs.Error05(false)].Zhi,

}
c.ServeJSON()
return
}
serviceret:=zdjuesequanxianyewus.Shanchujuesequanxian(id)
tishi:=chushihuas.Tishis[serviceret].Zhi
if tishi==zfzhi.Zhi.Kzf(){
splitret:=strings.Split(serviceret,zfzhi.Zhi.Xhx())
c.Data[zf.Zfs.Json(true)]=chushihuas.Tishis[splitret[zfzhi.Zhi.Shuzi0()]].Zhi+zfzhi.Zhi.Mh()+splitret[zfzhi.Zhi.Shuzi1()]
c.ServeJSON()
return
}
c.Data[zf.Zfs.Json(true)]=tishi
c.ServeJSON()
return
}
