package zdjuesekus
import (
"mhsydata/moxings"
"mhsydata/chushihuas"
"changliang/zf"
)
func Chaxunyige(id int) *moxings.Juese {
juese := &moxings.Juese{Id:id}
err:=chushihuas.Defaultormer().Read(juese)
if err !=nil{
return nil
}
return juese
}
func Tianjiayige(juese *moxings.Juese)string{
_,err:=chushihuas.Defaultormer().Insert(juese)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
}
return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Tianjiaduoge(jueseshuzu []moxings.Juese)string{
_,err:=chushihuas.Defaultormer().InsertMulti(len(jueseshuzu),jueseshuzu)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi04(false)].Bianma
}
return chushihuas.Tishis[zf.Zfs.Tishi03(false)].Bianma
}
func Xiugaiyige(juese *moxings.Juese)string{
_,err:=chushihuas.Defaultormer().Update(juese)
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
func Chaxunquanbu()[]*moxings.Juese{
ret:=[]*moxings.Juese{}
queryseter:=chushihuas.Defaultormer().QueryTable(zf.Zfs.Juese(true))
queryseter.All(&ret)
return ret
}
