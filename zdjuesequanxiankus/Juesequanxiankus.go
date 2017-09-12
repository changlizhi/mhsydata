package zdjuesequanxiankus
import (
"mhsydata/moxings"
"mhsydata/chushihuas"
"changliang/zf"
)
func Chaxunyige(id int) *moxings.Juesequanxian {
juesequanxian := &moxings.Juesequanxian{Id:id}
err:=chushihuas.Defaultormer().Read(juesequanxian)
if err !=nil{
return nil
}
return juesequanxian
}
func Tianjiayige(juesequanxian *moxings.Juesequanxian)string{
_,err:=chushihuas.Defaultormer().Insert(juesequanxian)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
}
return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(juesequanxianshuzu []moxings.Juesequanxian)string{
_,err:=chushihuas.Defaultormer().InsertMulti(len(juesequanxianshuzu),juesequanxianshuzu)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
}
return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Xiugaiyige(juesequanxian *moxings.Juesequanxian)string{
_,err:=chushihuas.Defaultormer().Update(juesequanxian)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi06(false)].Bianma
}
return chushihuas.Tishis[zf.Zfs.Tishi05(false)].Bianma
}
func Shanchuyige(id int)string{
_,err:=chushihuas.Defaultormer().Delete(Chaxunyige(id))
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi08(false)].Bianma
}
return chushihuas.Tishis[zf.Zfs.Tishi07(false)].Bianma
}
