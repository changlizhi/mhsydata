package tests
import(
"mhsydata/moxings"
"testing"
"log"
"changliang/zfzhi"
"mhsydata/zdjuesequanxiankus"
)
func TestTianjiaduogeJuesequanxian(t *testing.T){
juesequanxian2:=moxings.Juesequanxian{
Biaoji:"BiaojiTianjiaduogeTest2",
Id:zfzhi.Zhi.Shuzi1(),
Juesebianma:"JuesebianmaTianjiaduogeTest2",
Quanxianbianma:"QuanxianbianmaTianjiaduogeTest2",
}
juesequanxian3:=moxings.Juesequanxian{
Juesebianma:"JuesebianmaTianjiaduogeTest3",
Quanxianbianma:"QuanxianbianmaTianjiaduogeTest3",
Biaoji:"BiaojiTianjiaduogeTest3",
Id:zfzhi.Zhi.Shuzi1(),
}
juesequanxians:=[]moxings.Juesequanxian{juesequanxian2,juesequanxian3}
zdjuesequanxiankus.Tianjiaduoge(juesequanxians)
}
func TestTianjiayigeJuesequanxian(t *testing.T){
juesequanxian:=&moxings.Juesequanxian{
Id:zfzhi.Zhi.Shuzi1(),
Juesebianma:"JuesebianmaTianjiayigeTest1",
Quanxianbianma:"QuanxianbianmaTianjiayigeTest1",
Biaoji:"BiaojiTianjiayigeTest1",
}
zdjuesequanxiankus.Tianjiayige(juesequanxian)
}
func TestXiugaiyigeJuesequanxian(t *testing.T){
juesequanxian:=&moxings.Juesequanxian{
Biaoji:"BiaojiXiugaiyigeTest1",
Id:zfzhi.Zhi.Shuzi1(),
Juesebianma:"JuesebianmaXiugaiyigeTest1",
Quanxianbianma:"QuanxianbianmaXiugaiyigeTest1",
}
zdjuesequanxiankus.Xiugaiyige(juesequanxian)
}
func TestChaxunyigeJuesequanxian(t *testing.T){
juesequanxian:=zdjuesequanxiankus.Chaxunyige(zfzhi.Zhi.Shuzi1())
log.Println(juesequanxian)
}
func TestShanchuyigeJuesequanxian(t *testing.T){
zdjuesequanxiankus.Shanchuyige(zfzhi.Zhi.Shuzi1())
}
