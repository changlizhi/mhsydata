package zdkongzhiqis
import(
"changliang/zf"
"github.com/astaxie/beego"

"gongju"

)
type Mainkongzhiqi struct{
beego.Controller
}
func (c *Mainkongzhiqi)Get(){
c.Data[zf.Zfs.Json(true)]=gongju.Mokuaimings
c.ServeJSON()

}
