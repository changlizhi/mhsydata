package zddtziyuankus

import (
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"mhsydata/chushihuas"
	"mhsydata/moxings"
)

func Chaxunyigekus(id int) *moxings.Dtziyuan {
	dtziyuan := &moxings.Dtziyuan{Id: id}
	err := chushihuas.Defaultormer().Read(dtziyuan)
	if err != nil {
		log.Println(err)
		return nil
	}
	return dtziyuan
}
func Tianjiayigekus(dtziyuan *moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().Insert(dtziyuan)
	if err != nil {
		log.Println(err)
		return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduogekus(dtziyuanshuzu []*moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().InsertMulti(len(dtziyuanshuzu), dtziyuanshuzu)
	if err != nil {
		log.Println(err)
		return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Xiugaiyigekus(dtziyuan *moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().Update(dtziyuan)
	if err != nil {
		log.Println(err)
		return chushihuas.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
func Shanchuyigekus(id int) string {
	_, err := chushihuas.Defaultormer().Delete(Chaxunyigekus(id))
	if err != nil {
		log.Println(err)
		return chushihuas.Tishis[zf.Zfs.Tishi08(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi07(false)].Bianma
}
func Chaxunquanbukus(dtziyuan *moxings.Dtziyuan) []*moxings.Dtziyuan {
	ret := []*moxings.Dtziyuan{}
	queryseter := chushihuas.Defaultormer().QueryTable(zf.Zfs.Dtziyuan(true))
	queryseter = queryseter.Limit(zfzhi.Zhi.Shuzifu1())
	queryseter = queryseter.OrderBy(zf.Zfs.Paixu(true))
	houzhui := zfzhi.Zhi.Xhx() + zfzhi.Zhi.Xhx() + zf.Zfs.I(true) + zf.Zfs.Exact(true)
	if dtziyuan.Fubianma != zfzhi.Zhi.Kzf() {
		queryseter = queryseter.Filter(zf.Zfs.Fubianma(true)+houzhui, dtziyuan.Fubianma)

	}
	if dtziyuan.Lujing != zfzhi.Zhi.Kzf() {
		queryseter = queryseter.Filter(zf.Zfs.Lujing(true)+houzhui, dtziyuan.Lujing)

	}
	if dtziyuan.Mingcheng != zfzhi.Zhi.Kzf() {
		queryseter = queryseter.Filter(zf.Zfs.Mingcheng(true)+houzhui, dtziyuan.Mingcheng)

	}
	if dtziyuan.Bianma != zfzhi.Zhi.Kzf() {
		queryseter = queryseter.Filter(zf.Zfs.Bianma(true)+houzhui, dtziyuan.Bianma)

	}
	if dtziyuan.Biaoji != zfzhi.Zhi.Kzf() {
		queryseter = queryseter.Filter(zf.Zfs.Biaoji(true)+houzhui, dtziyuan.Biaoji)

	}
	_, err := queryseter.All(&ret)
	if err != nil {
		log.Println(err)
		return nil
	}
	return ret
}
