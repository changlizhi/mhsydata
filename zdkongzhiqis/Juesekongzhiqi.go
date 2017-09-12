package zdkongzhiqis
import(
"encoding/json"
"github.com/astaxie/beego"
"mhsydata/moxings"
"mhsydata/chushihuas"
"mhsydata/zdjueseyewus"
"changliang/zf"
"changliang/zfzhi"
"log"
"strconv"
"strings"

)
type Juesekongzhiqi struct{
beego.Controller
}
func (c *Juesekongzhiqi)Get(){
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
juese:=zdjueseyewus.Chaxunjuese(id)
c.Data[zf.Zfs.Json(true)]=juese
c.ServeJSON()
return
}
func (c *Juesekongzhiqi)Post(){
juese:=moxings.Juese{}
json.Unmarshal(c.Ctx.Input.RequestBody,&juese)
serviceret:=zdjueseyewus.Tianjiajuese(&juese)
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
func (c *Juesekongzhiqi)Patch(){
juese:=moxings.Juese{}
json.Unmarshal(c.Ctx.Input.RequestBody,&juese)
serviceret:=zdjueseyewus.Xiugaijuese(&juese)
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
func (c *Juesekongzhiqi)Delete(){
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
serviceret:=zdjueseyewus.Shanchujuese(id)
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
