package tests
import(
"mhsydata/moxings"
"testing"
"log"
"changliang/zfzhi"
"mhsydata/zdjuesekus"
)
func TestTianjiaduogeJuese(t *testing.T){
juese2:=moxings.Juese{
Bianma:"BianmaTianjiaduogeTest2",
Mingcheng:"MingchengTianjiaduogeTest2",
Id:zfzhi.Zhi.Shuzi1(),
Biaoji:"BiaojiTianjiaduogeTest2",
}
juese3:=moxings.Juese{
Bianma:"BianmaTianjiaduogeTest3",
Mingcheng:"MingchengTianjiaduogeTest3",
Id:zfzhi.Zhi.Shuzi1(),
Biaoji:"BiaojiTianjiaduogeTest3",
}
jueses:=[]moxings.Juese{juese2,juese3}
zdjuesekus.Tianjiaduoge(jueses)
}
func TestTianjiayigeJuese(t *testing.T){
juese:=&moxings.Juese{
Id:zfzhi.Zhi.Shuzi1(),
Biaoji:"BiaojiTianjiayigeTest1",
Bianma:"BianmaTianjiayigeTest1",
Mingcheng:"MingchengTianjiayigeTest1",
}
zdjuesekus.Tianjiayige(juese)
}
func TestXiugaiyigeJuese(t *testing.T){
juese:=&moxings.Juese{
Id:zfzhi.Zhi.Shuzi1(),
Biaoji:"BiaojiXiugaiyigeTest1",
Bianma:"BianmaXiugaiyigeTest1",
Mingcheng:"MingchengXiugaiyigeTest1",
}
zdjuesekus.Xiugaiyige(juese)
}
func TestChaxunyigeJuese(t *testing.T){
juese:=zdjuesekus.Chaxunyige(zfzhi.Zhi.Shuzi1())
log.Println(juese)
}
func TestShanchuyigeJuese(t *testing.T){
zdjuesekus.Shanchuyige(zfzhi.Zhi.Shuzi1())
}
