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
Xiugairiqi:time.Now(),
Biaoji:"DtziyuanTianjia",
Id:zfzhi.Zhi.Shuzi1(),
Mingcheng:"DtziyuanTianjia",
Paixu:"DtziyuanTianjia",
Bianma:"DtziyuanTianjia",
Fubianma:"DtziyuanTianjia",
Lujing:"DtziyuanTianjia",

}
zddtziyuanyewus.Tianjiadtziyuan(&dtziyuan)
}

func TestDtziyuanyewusXiugai(t *testing.T){
dtziyuan:=moxings.Dtziyuan{
Id:zfzhi.Zhi.Shuzi1(),
Mingcheng:"DtziyuanXiugai",
Paixu:"DtziyuanXiugai",
Bianma:"DtziyuanXiugai",
Fubianma:"DtziyuanXiugai",
Lujing:"DtziyuanXiugai",
Xiugairiqi:time.Now(),
Biaoji:"DtziyuanXiugai",

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
