package zdjuesequanxianyewus
import(
"mhsydata/suoyoucuowus"
"bytes"
"time"
"mhsydata/moxings"
"gongju"
"changliang/zf"
"changliang/zfzhi"
"mhsydata/zdjuesequanxiankus"
"mhsydata/chushihuas"
)
func yanzhengziduanchangdu(juesequanxian *moxings.Juesequanxian)error{
cuowu:=false
buffer:=bytes.Buffer{}

lenbiaoji:=gongju.Liechangdu(zf.Zfs.Biaoji(false))
lenbiaojishiti:=len(juesequanxian.Biaoji)
if lenbiaojishiti>int(lenbiaoji){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenbiaoji),lenbiaojishiti))
}

lenjuesebianma:=gongju.Liechangdu(zf.Zfs.Juesebianma(false))
lenjuesebianmashiti:=len(juesequanxian.Juesebianma)
if lenjuesebianmashiti>int(lenjuesebianma){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenjuesebianma),lenjuesebianmashiti))
}

lenquanxianbianma:=gongju.Liechangdu(zf.Zfs.Quanxianbianma(false))
lenquanxianbianmashiti:=len(juesequanxian.Quanxianbianma)
if lenquanxianbianmashiti>int(lenquanxianbianma){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenquanxianbianma),lenquanxianbianmashiti))
}
if cuowu{
return suoyoucuowus.Ziduancuowu{Shijian:time.Now(),Wenti:buffer.String()}
}
return nil
}
func Tianjiajuesequanxian(juesequanxian *moxings.Juesequanxian)string{
err:=yanzhengziduanchangdu(juesequanxian)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi09(false)].Bianma+zfzhi.Zhi.Xhx()+err.Error()

}
return zdjuesequanxiankus.Tianjiayige(juesequanxian)

}
func Xiugaijuesequanxian(juesequanxian *moxings.Juesequanxian)string{
err:=yanzhengziduanchangdu(juesequanxian)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi09(false)].Bianma+zfzhi.Zhi.Xhx()+err.Error()

}
juesequanxianfind:=Chaxunjuesequanxian(juesequanxian.Id)
if juesequanxianfind!=nil{

if juesequanxian.Biaoji!=zfzhi.Zhi.Kzf(){
juesequanxianfind.Biaoji=juesequanxian.Biaoji
}
if juesequanxian.Juesebianma!=zfzhi.Zhi.Kzf(){
juesequanxianfind.Juesebianma=juesequanxian.Juesebianma
}
if juesequanxian.Quanxianbianma!=zfzhi.Zhi.Kzf(){
juesequanxianfind.Quanxianbianma=juesequanxian.Quanxianbianma
}
return zdjuesequanxiankus.Xiugaiyige(juesequanxianfind)

}
return chushihuas.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchujuesequanxian(id int)string{
return zdjuesequanxiankus.Shanchuyige(id)
}
func Chaxunjuesequanxian(id int)*moxings.Juesequanxian{
return zdjuesequanxiankus.Chaxunyige(id)
}
