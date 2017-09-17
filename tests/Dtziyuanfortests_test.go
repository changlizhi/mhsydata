package tests

import (
	"changliang/fanshe"
	"changliang/zfzhi"
	"log"
	"mhsydata/fortests"
	"testing"
)

func TestZuzhuangdtziyuan(t *testing.T) {
	obj := fortests.Zuzhuangdtziyuan("Luyou", zfzhi.Zhi.Shuzi4(), fanshe.Fangfaming(false))
	jstr := fortests.Zuzhuangdtziyuanstring("Luyou", zfzhi.Zhi.Shuzi4(), fanshe.Fangfaming(false))
	log.Println(obj)
	log.Println(jstr)
}
