package chushihuas

import (
	"changliang/zf"
	"changliang/zfzhi"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"mhsydata/moxings"
	"strconv"
)

func init() {
	ormdebug()
	ormermoxing()
	ormershujuku()

}
func ormdebug() {
	Setormerdebug(
		Chushihuas[zf.Zfs.Ormdebug(false)].Zhi,
	)

}
func ormermoxing() {
	orm.RegisterModel(
		new(moxings.Dtziyuan),
	)

}
func Setormerdebug(boolstr string) {
	orm.Debug, _ = strconv.ParseBool(boolstr)
}
func Defaultormer() orm.Ormer {
	return Ormerbyname(zf.Zfs.Default(true))
}
func Ormerbyname(name string) orm.Ormer {
	ret := orm.NewOrm()
	ret.Using(name)
	return ret
}
func ormershujuku() {
	url :=
		Shujukus[zf.Zfs.Yonghu(false)].Zhi +
			zfzhi.Zhi.Mh() +
			Shujukus[zf.Zfs.Mima(false)].Zhi +
			zfzhi.Zhi.Qa() +
			zf.Zfs.Tcp(true) +
			zfzhi.Zhi.Xkhz() +
			Shujukus[zf.Zfs.Ip(false)].Zhi +
			zfzhi.Zhi.Mh() +
			Shujukus[zf.Zfs.Duankou(false)].Zhi +
			zfzhi.Zhi.Xkhy() +
			zfzhi.Zhi.Xx() +
			Shujukus[zf.Zfs.Mingcheng(false)].Zhi
	orm.RegisterDataBase(
		zf.Zfs.Default(true),
		Shujukus[zf.Zfs.Qudong(false)].Zhi,
		url,
	)

}
