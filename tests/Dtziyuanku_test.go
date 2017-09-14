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
		Xiugairiqi:     time.Now(),
		Biaoji:         "BiaojiTianjiaduogeTest2",
		Mingcheng:      "MingchengTianjiaduogeTest2",
		Paixu:          zfzhi.Zhi.Shuzi1(),
		Bianma:         "BianmaTianjiaduogeTest2",
		Lujing:         "LujingTianjiaduogeTest2",
		Chuangjianriqi: time.Now(),
		Id:             zfzhi.Zhi.Shuzi1(),
		Fubianma:       "FubianmaTianjiaduogeTest2",
	}
	dtziyuan3 := moxings.Dtziyuan{
		Paixu:          zfzhi.Zhi.Shuzi1(),
		Chuangjianriqi: time.Now(),
		Lujing:         "LujingTianjiaduogeTest3",
		Xiugairiqi:     time.Now(),
		Biaoji:         "BiaojiTianjiaduogeTest3",
		Id:             zfzhi.Zhi.Shuzi1(),
		Mingcheng:      "MingchengTianjiaduogeTest3",
		Bianma:         "BianmaTianjiaduogeTest3",
		Fubianma:       "FubianmaTianjiaduogeTest3",
	}
	dtziyuans := []moxings.Dtziyuan{dtziyuan2, dtziyuan3}
	zddtziyuankus.Tianjiaduoge(dtziyuans)
}
func TestTianjiayigeDtziyuan(t *testing.T) {
	dtziyuan := &moxings.Dtziyuan{
		Id:             zfzhi.Zhi.Shuzi1(),
		Paixu:          zfzhi.Zhi.Shuzi1(),
		Bianma:         "BianmaTianjiayigeTest1",
		Chuangjianriqi: time.Now(),
		Mingcheng:      "MingchengTianjiayigeTest1",
		Fubianma:       "FubianmaTianjiayigeTest1",
		Lujing:         "LujingTianjiayigeTest1",
		Xiugairiqi:     time.Now(),
		Biaoji:         "BiaojiTianjiayigeTest1",
	}
	zddtziyuankus.Tianjiayige(dtziyuan)
}
func TestXiugaiyigeDtziyuan(t *testing.T) {
	dtziyuan := &moxings.Dtziyuan{
		Paixu:          zfzhi.Zhi.Shuzi1(),
		Chuangjianriqi: time.Now(),
		Fubianma:       "FubianmaXiugaiyigeTest1",
		Lujing:         "LujingXiugaiyigeTest1",
		Xiugairiqi:     time.Now(),
		Biaoji:         "BiaojiXiugaiyigeTest1",
		Id:             zfzhi.Zhi.Shuzi1(),
		Mingcheng:      "MingchengXiugaiyigeTest1",
		Bianma:         "BianmaXiugaiyigeTest1",
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
