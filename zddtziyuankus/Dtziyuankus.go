package zddtziyuankus

import (
	"changliang/zf"
	"mhsydata/chushihuas"
	"mhsydata/moxings"
)

func Chaxunyige(id int) *moxings.Dtziyuan {
	dtziyuan := &moxings.Dtziyuan{Id: id}
	err := chushihuas.Defaultormer().Read(dtziyuan)
	if err != nil {
		return nil
	}
	return dtziyuan
}
func Tianjiayige(dtziyuan *moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().Insert(dtziyuan)
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(dtziyuanshuzu []*moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().InsertMulti(len(dtziyuanshuzu), dtziyuanshuzu)
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Xiugaiyige(dtziyuan *moxings.Dtziyuan) string {
	_, err := chushihuas.Defaultormer().Update(dtziyuan)
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi06(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
func Shanchuyige(id int) string {
	_, err := chushihuas.Defaultormer().Delete(Chaxunyige(id))
	if err != nil {
		return chushihuas.Tishis[zf.Zfs.Tishi08(false)].Bianma
	}
	return chushihuas.Tishis[zf.Zfs.Tishi07(false)].Bianma
}
func Chaxunquanbu() []*moxings.Dtziyuan {
	ret := []*moxings.Dtziyuan{}
	queryseter := chushihuas.Defaultormer().QueryTable(zf.Zfs.Dtziyuan(true))
	queryseter.All(&ret)
	return ret
}
