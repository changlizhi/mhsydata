package zddtziyuanyewus

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"gongju"
	"mhsydata/chushihuas"
	"mhsydata/moxings"
	"mhsydata/suoyoucuowus"
	"mhsydata/zddtziyuankus"
	"time"
)

func yanzhengziduanchangdu(dtziyuan *moxings.Dtziyuan) error {
	cuowu := false
	buffer := bytes.Buffer{}

	lenlujing := gongju.Liechangdu(zf.Zfs.Lujing(false))
	lenlujingshiti := len(dtziyuan.Lujing)
	if lenlujingshiti > int(lenlujing) {
		cuowu = true
		buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenlujing), lenlujingshiti))
	}

	lenbianma := gongju.Liechangdu(zf.Zfs.Bianma(false))
	lenbianmashiti := len(dtziyuan.Bianma)
	if lenbianmashiti > int(lenbianma) {
		cuowu = true
		buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbianma), lenbianmashiti))
	}

	lenfubianma := gongju.Liechangdu(zf.Zfs.Fubianma(false))
	lenfubianmashiti := len(dtziyuan.Fubianma)
	if lenfubianmashiti > int(lenfubianma) {
		cuowu = true
		buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenfubianma), lenfubianmashiti))
	}

	lenbiaoji := gongju.Liechangdu(zf.Zfs.Biaoji(false))
	lenbiaojishiti := len(dtziyuan.Biaoji)
	if lenbiaojishiti > int(lenbiaoji) {
		cuowu = true
		buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenbiaoji), lenbiaojishiti))
	}

	lenmingcheng := gongju.Liechangdu(zf.Zfs.Mingcheng(false))
	lenmingchengshiti := len(dtziyuan.Mingcheng)
	if lenmingchengshiti > int(lenmingcheng) {
		cuowu = true
		buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false), int64(lenmingcheng), lenmingchengshiti))
	}
	if cuowu {
		return suoyoucuowus.Ziduancuowu{Shijian: time.Now(), Wenti: buffer.String()}
	}
	return nil
}
func Tianjiayigeyewus(dtziyuan *moxings.Dtziyuan) string {
	err := yanzhengziduanchangdu(dtziyuan)
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	return zddtziyuankus.Tianjiayigekus(dtziyuan)

}
func Xiugaiyigeyewus(dtziyuan *moxings.Dtziyuan) string {
	err := yanzhengziduanchangdu(dtziyuan)
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi09(false)].Bianma + zfzhi.Zhi.Xhx() + err.Error()

	}
	dtziyuanfind := Chaxunyigeyewus(dtziyuan.Id)
	if dtziyuanfind != nil {

		if dtziyuan.Biaoji != zfzhi.Zhi.Kzf() {
			dtziyuanfind.Biaoji = dtziyuan.Biaoji
		}
		if dtziyuan.Mingcheng != zfzhi.Zhi.Kzf() {
			dtziyuanfind.Mingcheng = dtziyuan.Mingcheng
		}
		if dtziyuan.Fubianma != zfzhi.Zhi.Kzf() {
			dtziyuanfind.Fubianma = dtziyuan.Fubianma
		}
		if dtziyuan.Bianma != zfzhi.Zhi.Kzf() {
			dtziyuanfind.Bianma = dtziyuan.Bianma
		}
		if dtziyuan.Lujing != zfzhi.Zhi.Kzf() {
			dtziyuanfind.Lujing = dtziyuan.Lujing
		}
		return zddtziyuankus.Xiugaiyigekus(dtziyuanfind)

	}
	return chushihuas.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchuyigeyewus(id int) string {
	return zddtziyuankus.Shanchuyigekus(id)
}
func Chaxunyigeyewus(id int) *moxings.Dtziyuan {
	return zddtziyuankus.Chaxunyigekus(id)
}
func Chaxunquanbuyewus() []*moxings.Dtziyuan {
	return zddtziyuankus.Chaxunquanbukus()
}
