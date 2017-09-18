package tests

import (
	"changliang/fanshe"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"mhsydata/fortests"
	"mhsydata/zddtziyuanyewus"
	"testing"
)

func TestTianjiayigeyewus(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Kus(false),
		zfzhi.Zhi.Shuzi1(),
		fanshe.Fangfaming(false),
	)
	zddtziyuanyewus.Tianjiayigeyewus(dtziyuan[zfzhi.Zhi.Shuzi0()])

}

func TestXiugaiyigeyewus(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Kus(false),
		zfzhi.Zhi.Shuzi1(),
		fanshe.Fangfaming(false),
	)
	zddtziyuanyewus.Xiugaiyigeyewus(dtziyuan[zfzhi.Zhi.Shuzi0()])

}
func TestChaxunyigeyewus(t *testing.T) {
	dtziyuan := zddtziyuanyewus.Chaxunyigeyewus(zfzhi.Zhi.Shuzi1())
	log.Println(dtziyuan)

}
func TestShanchuyigeyewus(t *testing.T) {
	zddtziyuanyewus.Shanchuyigeyewus(zfzhi.Zhi.Shuzi1())
}
func TestChaxunquanbuyewus(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Test(true),
		zfzhi.Zhi.Shuzi1(),
		fanshe.Fangfaming(false),
	)
	quanbu := zddtziyuanyewus.Chaxunquanbuyewus(dtziyuan[zfzhi.Zhi.Shuzi0()])
	log.Println("quanbu[zfzhi.Zhi.Shuzi0()]:====", quanbu[zfzhi.Zhi.Shuzi0()])
}
