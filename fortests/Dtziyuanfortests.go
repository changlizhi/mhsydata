package fortests

import (
	"changliang/zf"
	"changliang/zfzhi"
	"encoding/json"
	"mhsydata/moxings"
	"strconv"
	"time"
)

func Zuzhuangdtziyuan(leixing string, jige int, fangfa string) []*moxings.Dtziyuan {
	ret := make(
		[]*moxings.Dtziyuan,
		zfzhi.Zhi.Shuzi0(),
	)
	for i := zfzhi.Zhi.Shuzi1(); i <= jige; i++ {
		istring := strconv.Itoa(i)
		dtziyuan := &moxings.Dtziyuan{
			Chuangjianriqi: time.Now(),
			Xiugairiqi:     time.Now(),
			Biaoji:         leixing + zf.Zfs.Biaoji(false) + fangfa + istring,
			Id:             i,
			Mingcheng:      leixing + zf.Zfs.Mingcheng(false) + fangfa + istring,
			Paixu:          i,
			Bianma:         leixing + zf.Zfs.Bianma(false) + fangfa + istring,
			Fubianma:       leixing + zf.Zfs.Fubianma(false) + fangfa + istring,
			Lujing:         leixing + zf.Zfs.Lujing(false) + fangfa + istring,
		}
		ret = append(ret, dtziyuan)

	}
	return ret
}
func Zuzhuangdtziyuanstring(leixing string, jige int, fangfa string) string {
	ret := Zuzhuangdtziyuan(leixing, jige, fangfa)
	jstring, _ := json.Marshal(ret)
	return string(jstring)

}
