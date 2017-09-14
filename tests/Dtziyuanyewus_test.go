package tests
import (
"mhsydata/moxings"
"changliang/zfzhi"
"mhsydata/zddtziyuanyewus"
"time"
"log"
"testing"

)
func TestDtziyuanyewusTianjia(t *testing.T){
dtziyuan:=moxings.Dtziyuan{
Mingcheng:"DtziyuanTianjia",
Paixu:"DtziyuanTianjia",
Bianma:"DtziyuanTianjia",
Id:zfzhi.Zhi.Shuzi1(),
Lujing:"DtziyuanTianjia",
Chuangjianriqi:time.Now(),
Xiugairiqi:time.Now(),
Biaoji:"DtziyuanTianjia",
Fubianma:"DtziyuanTianjia",

}
zddtziyuanyewus.Tianjiadtziyuan(&dtziyuan)
}

func TestDtziyuanyewusXiugai(t *testing.T){
dtziyuan:=moxings.Dtziyuan{
Fubianma:"DtziyuanXiugai",
Xiugairiqi:time.Now(),
Biaoji:"DtziyuanXiugai",
Id:zfzhi.Zhi.Shuzi1(),
Mingcheng:"DtziyuanXiugai",
Paixu:"DtziyuanXiugai",
Bianma:"DtziyuanXiugai",
Lujing:"DtziyuanXiugai",
Chuangjianriqi:time.Now(),

}
zddtziyuanyewus.Xiugaidtziyuan(&dtziyuan)
}
func TestDtziyuanyewusChaxun(t *testing.T){
dtziyuan:=zddtziyuanyewus.Chaxundtziyuan(zfzhi.Zhi.Shuzi1())
log.Println(dtziyuan)

}
func TestDtziyuanyewusShanchu(t *testing.T){
zddtziyuanyewus.Shanchudtziyuan(zfzhi.Zhi.Shuzi1())
}
func TestDtziyuanyewusChaxunquanbu(t *testing.T){
quanbu:=zddtziyuanyewus.Chaxunquanbudtziyuan()
log.Println("quanbu[zfzhi.Zhi.Shuzi0()]:====",quanbu[zfzhi.Zhi.Shuzi0()])
}
