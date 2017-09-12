package tests
import(
"github.com/astaxie/beego/context"
"mhsydata/zdkongzhiqis"
"gongju"
"changliang/zf"
"changliang/zfzhi"
"log"
"strconv"
"testing"

)
func mainkongzhiqi()*zdkongzhiqis.Mainkongzhiqi{
c:=zdkongzhiqis.Mainkongzhiqi{}
c.Data=make(map[interface{}]interface{})
c.Ctx=context.NewContext()
c.Ctx.Input=context.NewInput()
c.Ctx.Output=context.NewOutput()
c.Ctx.Output.Context=context.NewContext()
c.Ctx.Output.Context.ResponseWriter=&context.Response{new(gongju.Mywriter),true,200}
return &c
}
func TestGetmain(t *testing.T){
paramid:=strconv.Itoa(zfzhi.Zhi.Shuzi1())
c:=mainkongzhiqi()
c.Ctx.Input.SetParam(zfzhi.Zhi.Mh()+zf.Zfs.Id(false),paramid)
c.Get()
log.Println(c.Data[zf.Zfs.Json(true)])

}
