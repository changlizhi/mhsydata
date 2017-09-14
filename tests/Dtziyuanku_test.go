package tests
import(
"time"
"mhsydata/moxings"
"testing"
"log"
"changliang/zfzhi"
"mhsydata/zddtziyuankus"
)
func TestTianjiaduogeDtziyuan(t *testing.T){
dtziyuan2:=moxings.Dtziyuan{
Xiugairiqi:time.Now(),
Biaoji:"BiaojiTianjiaduogeTest2",
Id:zfzhi.Zhi.Shuzi1(),
Mingcheng:"MingchengTianjiaduogeTest2",
Paixu:"PaixuTianjiaduogeTest2",
Bianma:"BianmaTianjiaduogeTest2",
Fubianma:"FubianmaTianjiaduogeTest2",
Lujing:"LujingTianjiaduogeTest2",
}
dtziyuan3:=moxings.Dtziyuan{
Paixu:"PaixuTianjiaduogeTest3",
Bianma:"BianmaTianjiaduogeTest3",
Fubianma:"FubianmaTianjiaduogeTest3",
Lujing:"LujingTianjiaduogeTest3",
Xiugairiqi:time.Now(),
Biaoji:"BiaojiTianjiaduogeTest3",
Id:zfzhi.Zhi.Shuzi1(),
Mingcheng:"MingchengTianjiaduogeTest3",
}
dtziyuans:=[]moxings.Dtziyuan{dtziyuan2,dtziyuan3}
zddtziyuankus.Tianjiaduoge(dtziyuans)
}
func TestTianjiayigeDtziyuan(t *testing.T){
dtziyuan:=&moxings.Dtziyuan{
Xiugairiqi:time.Now(),
Biaoji:"BiaojiTianjiayigeTest1",
Id:zfzhi.Zhi.Shuzi1(),
Mingcheng:"MingchengTianjiayigeTest1",
Paixu:"PaixuTianjiayigeTest1",
Bianma:"BianmaTianjiayigeTest1",
Fubianma:"FubianmaTianjiayigeTest1",
Lujing:"LujingTianjiayigeTest1",
}
zddtziyuankus.Tianjiayige(dtziyuan)
}
func TestXiugaiyigeDtziyuan(t *testing.T){
dtziyuan:=&moxings.Dtziyuan{
Bianma:"BianmaXiugaiyigeTest1",
Fubianma:"FubianmaXiugaiyigeTest1",
Lujing:"LujingXiugaiyigeTest1",
Xiugairiqi:time.Now(),
Biaoji:"BiaojiXiugaiyigeTest1",
Id:zfzhi.Zhi.Shuzi1(),
Mingcheng:"MingchengXiugaiyigeTest1",
Paixu:"PaixuXiugaiyigeTest1",
}
zddtziyuankus.Xiugaiyige(dtziyuan)
}
func TestChaxunyigeDtziyuan(t *testing.T){
dtziyuan:=zddtziyuankus.Chaxunyige(zfzhi.Zhi.Shuzi1())
log.Println(dtziyuan)
}
func TestShanchuyigeDtziyuan(t *testing.T){
zddtziyuankus.Shanchuyige(zfzhi.Zhi.Shuzi1())
}
func TestChaxunquanbuDtziyuan(t *testing.T){
all:=zddtziyuankus.Chaxunquanbu()
log.Println("all[zfzhi.Zhi.Shuzi0()]:====",all[zfzhi.Zhi.Shuzi0()])
}
