package tests
import(
"net/http"
"changliang/zf"
"changliang/zfzhi"
"net/http/httptest"
"github.com/astaxie/beego"

"bytes"
"testing"
"log"

)
func TestDtziyuanget(t *testing.T){
r,_:=http.NewRequest(
zf.Zfs.GET(false),
zfzhi.Zhi.Xx()+zf.Zfs.Dtziyuan(true)+zfzhi.Zhi.Xx()+zfzhi.Zhi.Shuzi1w(),
nil,

)
w:=httptest.NewRecorder()
beego.BeeApp.Handlers.ServeHTTP(w,r)
log.Println(w.Body.String())

}
func TestDtziyuandelete(t *testing.T){
r,_:=http.NewRequest(
zf.Zfs.DELETE(false),
zfzhi.Zhi.Xx()+zf.Zfs.Dtziyuan(true)+zfzhi.Zhi.Xx()+zfzhi.Zhi.Shuzi1w(),
nil,

)
w:=httptest.NewRecorder()
beego.BeeApp.Handlers.ServeHTTP(w,r)
log.Println(w.Body.String())

}
func TestDtziyuanpost(t *testing.T){
canshu:=zfzhi.Zhi.Postdtziyuan()
r,_:=http.NewRequest(
zf.Zfs.POST(false),
zfzhi.Zhi.Xx()+zf.Zfs.Dtziyuan(true),
bytes.NewBuffer([]byte(canshu)),

)
r.Header.Set(
zf.Zfs.Content(false)+zfzhi.Zhi.Jian()+zf.Zfs.Type(false),
zf.Zfs.Application(true)+zfzhi.Zhi.Xx()+zf.Zfs.Json(true),

)
w:=httptest.NewRecorder()
beego.BeeApp.Handlers.ServeHTTP(w,r)
log.Println(w.Body.String())

}
func TestDtziyuanpatch(t *testing.T){
canshu:=zfzhi.Zhi.Patchdtziyuan()
r,_:=http.NewRequest(
zf.Zfs.PATCH(false),
zfzhi.Zhi.Xx()+zf.Zfs.Dtziyuan(true),
bytes.NewBuffer([]byte(canshu)),

)
r.Header.Set(
zf.Zfs.Content(false)+zfzhi.Zhi.Jian()+zf.Zfs.Type(false),
zf.Zfs.Application(true)+zfzhi.Zhi.Xx()+zf.Zfs.Json(true),

)
w:=httptest.NewRecorder()
beego.BeeApp.Handlers.ServeHTTP(w,r)
log.Println(w.Body.String())

}
func TestDtziyuanliebiaopost(t *testing.T){
canshu:=zfzhi.Zhi.Postdtziyuan()
r,_:=http.NewRequest(
zf.Zfs.POST(false),
zfzhi.Zhi.Xx()+zf.Zfs.Dtziyuan(true)+zfzhi.Zhi.Xx()+zf.Zfs.Quanbu(true),
bytes.NewBuffer([]byte(canshu)),

)
r.Header.Set(
zf.Zfs.Content(false)+zfzhi.Zhi.Jian()+zf.Zfs.Type(false),
zf.Zfs.Application(true)+zfzhi.Zhi.Xx()+zf.Zfs.Json(true),

)
w:=httptest.NewRecorder()
beego.BeeApp.Handlers.ServeHTTP(w,r)
log.Println(w.Body.String())

}
