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
		Biaoji:    "JueseTianjia",
		Bianma:    "JueseTianjia",
		Mingcheng: "JueseTianjia",
		Id:        zfzhi.Zhi.Shuzi1(),
	}
	zdjueseyewus.Tianjiajuese(&juese)
}

func TestJueseyewusXiugai(t *testing.T) {
	juese := moxings.Juese{
		Biaoji:    "JueseXiugai",
		Bianma:    "JueseXiugai",
		Mingcheng: "JueseXiugai",
		Id:        zfzhi.Zhi.Shuzi1(),
	}
	zdjueseyewus.Xiugaijuese(&juese)
}
func TestJueseyewusChaxun(t *testing.T) {
	juese := zdjueseyewus.Chaxunjuese(1)
	log.Println(juese)

}
func TestJueseyewusShanchu(t *testing.T) {
	zdjueseyewus.Shanchujuese(1)
}
