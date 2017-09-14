package tests

import (
	"changliang/zfzhi"
	"log"
	"mhsydata/moxings"
	"mhsydata/zddtziyuanyewus"
	"testing"
	"time"
)

func TestDtziyuanyewusTianjia(t *testing.T) {
	dtziyuan := moxings.Dtziyuan{
		Id:             zfzhi.Zhi.Shuzi1(),
		Mingcheng:      "DtziyuanTianjia",
		Bianma:         "DtziyuanTianjia",
		Chuangjianriqi: time.Now(),
		Xiugairiqi:     time.Now(),
		Paixu:          zfzhi.Zhi.Shuzi1(),
		Fubianma:       "DtziyuanTianjia",
		Lujing:         "DtziyuanTianjia",
		Biaoji:         "DtziyuanTianjia",
	}
	zddtziyuanyewus.Tianjiadtziyuan(&dtziyuan)
}

func TestDtziyuanyewusXiugai(t *testing.T) {
	dtziyuan := moxings.Dtziyuan{
		Mingcheng:      "DtziyuanXiugai",
		Lujing:         "DtziyuanXiugai",
		Biaoji:         "DtziyuanXiugai",
		Xiugairiqi:     time.Now(),
		Id:             zfzhi.Zhi.Shuzi1(),
		Paixu:          zfzhi.Zhi.Shuzi1(),
		Bianma:         "DtziyuanXiugai",
		Fubianma:       "DtziyuanXiugai",
		Chuangjianriqi: time.Now(),
	}
	zddtziyuanyewus.Xiugaidtziyuan(&dtziyuan)
}
func TestDtziyuanyewusChaxun(t *testing.T) {
	dtziyuan := zddtziyuanyewus.Chaxundtziyuan(zfzhi.Zhi.Shuzi1())
	log.Println(dtziyuan)

}
func TestDtziyuanyewusShanchu(t *testing.T) {
	zddtziyuanyewus.Shanchudtziyuan(zfzhi.Zhi.Shuzi1())
}
func TestDtziyuanyewusChaxunquanbu(t *testing.T) {
	quanbu := zddtziyuanyewus.Chaxunquanbudtziyuan()
	log.Println("quanbu[zfzhi.Zhi.Shuzi0()]:====", quanbu[zfzhi.Zhi.Shuzi0()])
}
