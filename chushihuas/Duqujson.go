package chushihuas

import (
	"changliang/zf"
	"gongju"
)

func Shezhipath() string {
	return gongju.Getwenjianmulu(
		zf.Zfs.Mhsydata(true),
		zf.Zfs.Peizhi(true),
		zf.Zfs.Shezhi(false),
		zf.Zfs.Json(true),
	)

}
func Shezhijson() *Shezhi {
	shezhi := Shezhi{}
	ret := gongju.Jiexi(
		Shezhipath(),
		&shezhi,
	)
	return ret.(*Shezhi)
}
func Guojihuapath() string {
	return gongju.Getwenjianmulu(
		zf.Zfs.Mhsydata(true),
		zf.Zfs.Peizhi(true),
		Chushihuas[zf.Zfs.Yuyan(false)].Zhi,
		zf.Zfs.Json(true),
	)

}
func Guojihuajson() *Guojihua {
	guojihua := Guojihua{}
	ret := gongju.Jiexi(
		Guojihuapath(),
		&guojihua,
	)
	return ret.(*Guojihua)
}
