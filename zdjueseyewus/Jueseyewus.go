package zdjueseyewus
import(
"mhsydata/suoyoucuowus"
"bytes"
"time"
"mhsydata/moxings"
"gongju"
"changliang/zf"
"changliang/zfzhi"
"mhsydata/zdjuesekus"
"mhsydata/chushihuas"
)
func yanzhengziduanchangdu(juese *moxings.Juese)error{
cuowu:=false
buffer:=bytes.Buffer{}

lenmingcheng:=gongju.Liechangdu(zf.Zfs.Mingcheng(false))
lenmingchengshiti:=len(juese.Mingcheng)
if lenmingchengshiti>int(lenmingcheng){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenmingcheng),lenmingchengshiti))
}

lenbiaoji:=gongju.Liechangdu(zf.Zfs.Biaoji(false))
lenbiaojishiti:=len(juese.Biaoji)
if lenbiaojishiti>int(lenbiaoji){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenbiaoji),lenbiaojishiti))
}

lenbianma:=gongju.Liechangdu(zf.Zfs.Bianma(false))
lenbianmashiti:=len(juese.Bianma)
if lenbianmashiti>int(lenbianma){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenbianma),lenbianmashiti))
}
if cuowu{
return suoyoucuowus.Ziduancuowu{Shijian:time.Now(),Wenti:buffer.String()}
}
return nil
}
func Tianjiajuese(juese *moxings.Juese)string{
err:=yanzhengziduanchangdu(juese)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi09(false)].Bianma+zfzhi.Zhi.Xhx()+err.Error()

}
return zdjuesekus.Tianjiayige(juese)

}
func Xiugaijuese(juese *moxings.Juese)string{
err:=yanzhengziduanchangdu(juese)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi09(false)].Bianma+zfzhi.Zhi.Xhx()+err.Error()

}
juesefind:=Chaxunjuese(juese.Id)
if juesefind!=nil{

if juese.Biaoji!=zfzhi.Zhi.Kzf(){
juesefind.Biaoji=juese.Biaoji
}
if juese.Bianma!=zfzhi.Zhi.Kzf(){
juesefind.Bianma=juese.Bianma
}
if juese.Mingcheng!=zfzhi.Zhi.Kzf(){
juesefind.Mingcheng=juese.Mingcheng
}
return zdjuesekus.Xiugaiyige(juesefind)

}
return chushihuas.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchujuese(id int)string{
return zdjuesekus.Shanchuyige(id)
}
func Chaxunjuese(id int)*moxings.Juese{
return zdjuesekus.Chaxunyige(id)
}
