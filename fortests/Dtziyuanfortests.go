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
		istr := strconv.Itoa(i)
		dtziyuan := &moxings.Dtziyuan{
			Xiugairiqi:     time.Now(),
			Biaoji:         leixing + zf.Zfs.Biaoji(false) + fangfa + istr,
			Mingcheng:      leixing + zf.Zfs.Mingcheng(false) + fangfa + istr,
			Paixu:          i,
			Bianma:         leixing + zf.Zfs.Bianma(false) + fangfa + istr,
			Lujing:         leixing + zf.Zfs.Lujing(false) + fangfa + istr,
			Chuangjianriqi: time.Now(),
			Id:             i,
			Fubianma:       leixing + zf.Zfs.Fubianma(false) + fangfa + istr,
		}
		ret = append(ret, dtziyuan)
	}

	return ret
}
func Zuzhuangdtziyuanstr(leixing string, jige int, fangfa string) string {
	ret := Zuzhuangdtziyuan(leixing, jige, fangfa)
	jstr, _ := json.Marshal(ret)
	return string(jstr)
}
