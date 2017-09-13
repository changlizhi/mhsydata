package tests
import (
"mhsydata/moxings"
"changliang/zfzhi"
"mhsydata/zdjueseyewus"
"log"
"testing"

)
func TestJueseyewusTianjia(t *testing.T){
juese:=moxings.Juese{
Mingcheng:"JueseTianjia",
Id:zfzhi.Zhi.Shuzi1(),
Biaoji:"JueseTianjia",
Bianma:"JueseTianjia",

}
zdjueseyewus.Tianjiajuese(&juese)
}

func TestJueseyewusXiugai(t *testing.T){
juese:=moxings.Juese{
Mingcheng:"JueseXiugai",
Id:zfzhi.Zhi.Shuzi1(),
Biaoji:"JueseXiugai",
Bianma:"JueseXiugai",

}
zdjueseyewus.Xiugaijuese(&juese)
}
func TestJueseyewusChaxun(t *testing.T){
juese:=zdjueseyewus.Chaxunjuese(1)
log.Println(juese)

}
func TestJueseyewusShanchu(t *testing.T){
zdjueseyewus.Shanchujuese(1)
}
