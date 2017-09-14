package luyous
import(
"github.com/astaxie/beego"
"changliang/zf"
"changliang/zfzhi"
"mhsydata/zdkongzhiqis"

)
func init(){
beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Dtziyuan(true),&zdkongzhiqis.Dtziyuankongzhiqi{})
beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Dtziyuan(true)+zfzhi.Zhi.Xx()+zfzhi.Zhi.Mh()+zf.Zfs.Id(false),&zdkongzhiqis.Dtziyuankongzhiqi{})
beego.Router(zfzhi.Zhi.Xx()+zf.Zfs.Dtziyuan(true)+zfzhi.Zhi.Xx()+zf.Zfs.Quanbu(true),&zdkongzhiqis.Dtziyuanliebiaokongzhiqi{})

}
