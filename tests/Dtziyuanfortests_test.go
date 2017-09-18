package tests

import (
	"changliang/fanshe"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"mhsydata/fortests"
	"testing"
)

func TestZuzhuangdtziyuan(t *testing.T) {
	jiegouti := fortests.Zuzhuangdtziyuan(
		zf.Zfs.Test(true),
		zfzhi.Zhi.Shuzi4(),
		fanshe.Fangfaming(false),
	)
	log.Println("jiegouti:====", jiegouti)
	jiegoutiyigestring := fortests.Zuzhuangdtziyuanyigestring(
		zf.Zfs.Test(true),
		fanshe.Fangfaming(false),
	)
	log.Println("jiegoutiyigestring:====", jiegoutiyigestring)

}
