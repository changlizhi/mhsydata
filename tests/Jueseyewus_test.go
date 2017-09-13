package tests

import (
	"changliang/zfzhi"
	"log"
	"mhsydata/moxings"
	"mhsydata/zdjueseyewus"
	"testing"
)

func TestJueseyewusTianjia(t *testing.T) {
	juese := moxings.Juese{
		Mingcheng: "JueseTianjia",
		Id:        zfzhi.Zhi.Shuzi1(),
		Biaoji:    "JueseTianjia",
		Bianma:    "JueseTianjia",
	}
	zdjueseyewus.Tianjiajuese(&juese)
}

func TestJueseyewusXiugai(t *testing.T) {
	juese := moxings.Juese{
		Id:        zfzhi.Zhi.Shuzi1(),
		Biaoji:    "JueseXiugai",
		Bianma:    "JueseXiugai",
		Mingcheng: "JueseXiugai",
	}
	zdjueseyewus.Xiugaijuese(&juese)
}
func TestJueseyewusChaxun(t *testing.T) {
	juese := zdjueseyewus.Chaxunjuese(zfzhi.Zhi.Shuzi1())
	log.Println(juese)

}
func TestJueseyewusShanchu(t *testing.T) {
	zdjueseyewus.Shanchujuese(zfzhi.Zhi.Shuzi1())
}
func TestJueseyewusChaxunquanbu(t *testing.T) {
	quanbu := zdjueseyewus.Chaxunquanbujuese()
	log.Println("quanbu[zfzhi.Zhi.Shuzi0()]:====", quanbu[zfzhi.Zhi.Shuzi0()])
}
