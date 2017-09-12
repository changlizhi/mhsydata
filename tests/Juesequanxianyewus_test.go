package tests
import (
"mhsydata/moxings"
"changliang/zfzhi"
"mhsydata/zdjuesequanxianyewus"
"log"
"testing"

)
func TestJuesequanxianyewusTianjia(t *testing.T){
juesequanxian:=moxings.Juesequanxian{
Id:zfzhi.Zhi.Shuzi1(),
Juesebianma:"JuesequanxianTianjia",
Quanxianbianma:"JuesequanxianTianjia",
Biaoji:"JuesequanxianTianjia",

}
zdjuesequanxianyewus.Tianjiajuesequanxian(&juesequanxian)
}

func TestJuesequanxianyewusXiugai(t *testing.T){
juesequanxian:=moxings.Juesequanxian{
Biaoji:"JuesequanxianXiugai",
Id:zfzhi.Zhi.Shuzi1(),
Juesebianma:"JuesequanxianXiugai",
Quanxianbianma:"JuesequanxianXiugai",

}
zdjuesequanxianyewus.Xiugaijuesequanxian(&juesequanxian)
}
func TestJuesequanxianyewusChaxun(t *testing.T){
juesequanxian:=zdjuesequanxianyewus.Chaxunjuesequanxian(1)
log.Println(juesequanxian)

}
func TestJuesequanxianyewusShanchu(t *testing.T){
zdjuesequanxianyewus.Shanchujuesequanxian(1)
}
