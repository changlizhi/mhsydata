package zddtziyuankus

import (
	"changliang/zf"
	"mhsydata/chushihuas"
	"mhsydata/moxings"
)

func Chaxunyigekus(id int) *moxings.Dtziyuan {
	dtziyuan := &moxings.Dtziyuan{Id: id}
	err := chushihuas.Defaultormer().Read(dtziyuan)
	if err != nil {
		return nil
	}
	return dtziyuan
}
func Tianjiayigekus(dtziyuan *moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().Insert(dtziyuan)
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduogekus(dtziyuanshuzu []*moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().InsertMulti(len(dtziyuanshuzu), dtziyuanshuzu)
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Xiugaiyigekus(dtziyuan *moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().Update(dtziyuan)
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
func Shanchuyigekus(id int) string {
	_, err := chushihuas.Defaultormer().Delete(Chaxunyigekus(id))
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi08(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi07(false)].Bianma
}
func Chaxunquanbukus() []*moxings.Dtziyuan {
	ret := []*moxings.Dtziyuan{}
	queryseter := chushihuas.Defaultormer().QueryTable(zf.Zfs.Dtziyuan(true))
	queryseter.All(&ret)
	return ret
}
