package tests

import (
	"changliang/zfzhi"
	"log"
	"mhsydata/moxings"
	"mhsydata/zddtziyuanyewus"
	"testing"
	"time"
)

func TestTianjiayigeyewus(t *testing.T) {
	dtziyuan := moxings.Dtziyuan{
		Paixu:          zfzhi.Zhi.Shuzi1(),
		Biaoji:         "DtziyuanTianjiayige",
		Id:             zfzhi.Zhi.Shuzi1(),
		Mingcheng:      "DtziyuanTianjiayige",
		Bianma:         "DtziyuanTianjiayige",
		Fubianma:       "DtziyuanTianjiayige",
		Lujing:         "DtziyuanTianjiayige",
		Chuangjianriqi: time.Now(),
		Xiugairiqi:     time.Now(),
	}
	zddtziyuanyewus.Tianjiayigeyewus(&dtziyuan)
}

func TestXiugaiyigeyewus(t *testing.T) {
	dtziyuan := moxings.Dtziyuan{
		Biaoji:         "DtziyuanXiugaiyige",
		Id:             zfzhi.Zhi.Shuzi1(),
		Mingcheng:      "DtziyuanXiugaiyige",
		Bianma:         "DtziyuanXiugaiyige",
		Fubianma:       "DtziyuanXiugaiyige",
		Lujing:         "DtziyuanXiugaiyige",
		Chuangjianriqi: time.Now(),
		Paixu:          zfzhi.Zhi.Shuzi1(),
		Xiugairiqi:     time.Now(),
	}
	zddtziyuanyewus.Xiugaiyigeyewus(&dtziyuan)
}
func TestChaxunyigeyewus(t *testing.T) {
	dtziyuan := zddtziyuanyewus.Chaxunyigeyewus(zfzhi.Zhi.Shuzi1())
	log.Println(dtziyuan)

}
func TestShanchuyigeyewus(t *testing.T) {
	zddtziyuanyewus.Shanchuyigeyewus(zfzhi.Zhi.Shuzi1())
}
func TestChaxunquanbuyewus(t *testing.T) {
	quanbu := zddtziyuanyewus.Chaxunquanbuyewus()
	log.Println("quanbu[zfzhi.Zhi.Shuzi0()]:====", quanbu[zfzhi.Zhi.Shuzi0()])
}
