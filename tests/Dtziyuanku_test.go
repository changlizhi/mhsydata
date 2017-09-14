package tests

import (
	"changliang/zfzhi"
	"log"
	"mhsydata/moxings"
	"mhsydata/zddtziyuankus"
	"testing"
	"time"
)

func TestTianjiaduogeDtziyuan(t *testing.T) {
	dtziyuan2 := moxings.Dtziyuan{
		Mingcheng:      "MingchengTianjiaduogeTest2",
		Paixu:          "PaixuTianjiaduogeTest2",
		Bianma:         "BianmaTianjiaduogeTest2",
		Fubianma:       "FubianmaTianjiaduogeTest2",
		Xiugairiqi:     time.Now(),
		Id:             zfzhi.Zhi.Shuzi1(),
		Chuangjianriqi: time.Now(),
		Biaoji:         "BiaojiTianjiaduogeTest2",
		Lujing:         "LujingTianjiaduogeTest2",
	}
	dtziyuan3 := moxings.Dtziyuan{
		Id:             zfzhi.Zhi.Shuzi1(),
		Mingcheng:      "MingchengTianjiaduogeTest3",
		Paixu:          "PaixuTianjiaduogeTest3",
		Bianma:         "BianmaTianjiaduogeTest3",
		Fubianma:       "FubianmaTianjiaduogeTest3",
		Chuangjianriqi: time.Now(),
		Xiugairiqi:     time.Now(),
		Lujing:         "LujingTianjiaduogeTest3",
		Biaoji:         "BiaojiTianjiaduogeTest3",
	}
	dtziyuans := []moxings.Dtziyuan{dtziyuan2, dtziyuan3}
	zddtziyuankus.Tianjiaduoge(dtziyuans)
}
func TestTianjiayigeDtziyuan(t *testing.T) {
	dtziyuan := &moxings.Dtziyuan{
		Paixu:          "PaixuTianjiayigeTest1",
		Bianma:         "BianmaTianjiayigeTest1",
		Chuangjianriqi: time.Now(),
		Biaoji:         "BiaojiTianjiayigeTest1",
		Id:             zfzhi.Zhi.Shuzi1(),
		Mingcheng:      "MingchengTianjiayigeTest1",
		Fubianma:       "FubianmaTianjiayigeTest1",
		Lujing:         "LujingTianjiayigeTest1",
		Xiugairiqi:     time.Now(),
	}
	zddtziyuankus.Tianjiayige(dtziyuan)
}
func TestXiugaiyigeDtziyuan(t *testing.T) {
	dtziyuan := &moxings.Dtziyuan{
		Paixu:          "PaixuXiugaiyigeTest1",
		Id:             zfzhi.Zhi.Shuzi1(),
		Mingcheng:      "MingchengXiugaiyigeTest1",
		Bianma:         "BianmaXiugaiyigeTest1",
		Fubianma:       "FubianmaXiugaiyigeTest1",
		Lujing:         "LujingXiugaiyigeTest1",
		Chuangjianriqi: time.Now(),
		Xiugairiqi:     time.Now(),
		Biaoji:         "BiaojiXiugaiyigeTest1",
	}
	zddtziyuankus.Xiugaiyige(dtziyuan)
}
func TestChaxunyigeDtziyuan(t *testing.T) {
	dtziyuan := zddtziyuankus.Chaxunyige(zfzhi.Zhi.Shuzi1())
	log.Println(dtziyuan)
}
func TestShanchuyigeDtziyuan(t *testing.T) {
	zddtziyuankus.Shanchuyige(zfzhi.Zhi.Shuzi1())
}
func TestChaxunquanbuDtziyuan(t *testing.T) {
	all := zddtziyuankus.Chaxunquanbu()
	log.Println("all[zfzhi.Zhi.Shuzi0()]:====", all[zfzhi.Zhi.Shuzi0()])
}
