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

func TestTianjiaduogeDtziyuan(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Kus(false),
		zfzhi.Zhi.Shuzi3(),
		fanshe.Fangfaming(false),
	)
	zddtziyuankus.Tianjiaduoge(dtziyuan)

}
func TestTianjiayigeDtziyuan(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Kus(false),
		zfzhi.Zhi.Shuzi1(),
		fanshe.Fangfaming(false),
	)
	zddtziyuankus.Tianjiayige(dtziyuan[zfzhi.Zhi.Shuzi0()])

}
func TestXiugaiyigeDtziyuan(t *testing.T) {
	dtziyuan := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Kus(false),
		zfzhi.Zhi.Shuzi1(),
		fanshe.Fangfaming(false),
	)
	zddtziyuankus.Xiugaiyige(dtziyuan[zfzhi.Zhi.Shuzi0()])

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
