package tests

import (
	"changliang/fanshe"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"mhsydata/fortests"
	"mhsydata/zddtziyuankus"
	"testing"
)

func TestTianjiaduogeDtziyuankus(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Kus(false),
		zfzhi.Zhi.Shuzi3(),
		fanshe.Fangfaming(false),
	)
	zddtziyuankus.Tianjiaduogekus(dtziyuan)

}
func TestTianjiayigeDtziyuankus(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Kus(false),
		zfzhi.Zhi.Shuzi1(),
		fanshe.Fangfaming(false),
	)
	zddtziyuankus.Tianjiayigekus(dtziyuan[zfzhi.Zhi.Shuzi0()])

}
func TestXiugaiyigeDtziyuankus(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Kus(false),
		zfzhi.Zhi.Shuzi1(),
		fanshe.Fangfaming(false),
	)
	zddtziyuankus.Xiugaiyigekus(dtziyuan[zfzhi.Zhi.Shuzi0()])

}
func TestChaxunyigeDtziyuankus(t *testing.T) {
	dtziyuan := zddtziyuankus.Chaxunyigekus(zfzhi.Zhi.Shuzi1())
	log.Println(dtziyuan)
}
func TestShanchuyigeDtziyuan(t *testing.T) {
	zddtziyuankus.Shanchuyigekus(zfzhi.Zhi.Shuzi1())
}
func TestChaxunquanbuDtziyuankus(t *testing.T) {
	all := zddtziyuankus.Chaxunquanbukus()
	log.Println("all[zfzhi.Zhi.Shuzi0()]:====", all[zfzhi.Zhi.Shuzi0()])
}
