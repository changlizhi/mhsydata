package chushihuas

import (
	"changliang/zfzhi"
	"gongju"
)

func Getapppath() string {
	return gongju.Getpath(zfzhi.Zhi.Kzf())
}
