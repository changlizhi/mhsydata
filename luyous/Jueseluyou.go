package luyous
import(
"github.com/astaxie/beego"
"changliang/zf"
"changliang/zfzhi"
"mhsydata/zdkongzhiqis"

)
func init(){
beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Juese(true),&zdkongzhiqis.Juesekongzhiqi{})
beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Juese(true)+zfzhi.Zhi.Xx()+zfzhi.Zhi.Mh()+zf.Zfs.Id(false),&zdkongzhiqis.Juesekongzhiqi{})

}
