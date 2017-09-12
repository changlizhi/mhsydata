package luyous
import(
"github.com/astaxie/beego"
"changliang/zf"
"changliang/zfzhi"
"mhsydata/zdkongzhiqis"

)
func init(){
beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Juesequanxian(true),&zdkongzhiqis.Juesequanxiankongzhiqi{})
beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Juesequanxian(true)+zfzhi.Zhi.Xx()+zfzhi.Zhi.Mh()+zf.Zfs.Id(false),&zdkongzhiqis.Juesequanxiankongzhiqi{})

}
