package zddtziyuanyewus
import(
"mhsydata/suoyoucuowus"
"bytes"
"time"
"mhsydata/moxings"
"gongju"
"changliang/zf"
"changliang/zfzhi"
"mhsydata/zddtziyuankus"
"mhsydata/chushihuas"
)
func yanzhengziduanchangdu(dtziyuan *moxings.Dtziyuan)error{
cuowu:=false
buffer:=bytes.Buffer{}

lenpaixu:=gongju.Liechangdu(zf.Zfs.Paixu(false))
lenpaixushiti:=len(dtziyuan.Paixu)
if lenpaixushiti>int(lenpaixu){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenpaixu),lenpaixushiti))
}

lenfubianma:=gongju.Liechangdu(zf.Zfs.Fubianma(false))
lenfubianmashiti:=len(dtziyuan.Fubianma)
if lenfubianmashiti>int(lenfubianma){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenfubianma),lenfubianmashiti))
}

lenlujing:=gongju.Liechangdu(zf.Zfs.Lujing(false))
lenlujingshiti:=len(dtziyuan.Lujing)
if lenlujingshiti>int(lenlujing){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenlujing),lenlujingshiti))
}

lenbiaoji:=gongju.Liechangdu(zf.Zfs.Biaoji(false))
lenbiaojishiti:=len(dtziyuan.Biaoji)
if lenbiaojishiti>int(lenbiaoji){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenbiaoji),lenbiaojishiti))
}

lenmingcheng:=gongju.Liechangdu(zf.Zfs.Mingcheng(false))
lenmingchengshiti:=len(dtziyuan.Mingcheng)
if lenmingchengshiti>int(lenmingcheng){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenmingcheng),lenmingchengshiti))
}

lenbianma:=gongju.Liechangdu(zf.Zfs.Bianma(false))
lenbianmashiti:=len(dtziyuan.Bianma)
if lenbianmashiti>int(lenbianma){
cuowu=true
buffer.WriteString(gongju.Shengchengerrorchangdu(zf.Zfs.Mingcheng(false),int64(lenbianma),lenbianmashiti))
}
if cuowu{
return suoyoucuowus.Ziduancuowu{Shijian:time.Now(),Wenti:buffer.String()}
}
return nil
}
func Tianjiadtziyuan(dtziyuan *moxings.Dtziyuan)string{
err:=yanzhengziduanchangdu(dtziyuan)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi09(false)].Bianma+zfzhi.Zhi.Xhx()+err.Error()

}
return zddtziyuankus.Tianjiayige(dtziyuan)

}
func Xiugaidtziyuan(dtziyuan *moxings.Dtziyuan)string{
err:=yanzhengziduanchangdu(dtziyuan)
if err!=nil{
return chushihuas.Tishis[zf.Zfs.Tishi09(false)].Bianma+zfzhi.Zhi.Xhx()+err.Error()

}
dtziyuanfind:=Chaxundtziyuan(dtziyuan.Id)
if dtziyuanfind!=nil{

if dtziyuan.Mingcheng!=zfzhi.Zhi.Kzf(){
dtziyuanfind.Mingcheng=dtziyuan.Mingcheng
}
if dtziyuan.Lujing!=zfzhi.Zhi.Kzf(){
dtziyuanfind.Lujing=dtziyuan.Lujing
}
if dtziyuan.Biaoji!=zfzhi.Zhi.Kzf(){
dtziyuanfind.Biaoji=dtziyuan.Biaoji
}
if dtziyuan.Paixu!=zfzhi.Zhi.Kzf(){
dtziyuanfind.Paixu=dtziyuan.Paixu
}
if dtziyuan.Bianma!=zfzhi.Zhi.Kzf(){
dtziyuanfind.Bianma=dtziyuan.Bianma
}
if dtziyuan.Fubianma!=zfzhi.Zhi.Kzf(){
dtziyuanfind.Fubianma=dtziyuan.Fubianma
}
return zddtziyuankus.Xiugaiyige(dtziyuanfind)

}
return chushihuas.Cuowus[zf.Zfs.Error04(false)].Zhi
}
func Shanchudtziyuan(id int)string{
return zddtziyuankus.Shanchuyige(id)
}
func Chaxundtziyuan(id int)*moxings.Dtziyuan{
return zddtziyuankus.Chaxunyige(id)
}
func Chaxunquanbudtziyuan()[]*moxings.Dtziyuan{
return zddtziyuankus.Chaxunquanbu()
}
