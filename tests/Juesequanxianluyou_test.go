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
func TestJuesequanxianget(t *testing.T){
r,_:=http.NewRequest(
zf.Zfs.GET(false),
zfzhi.Zhi.Xx()+zf.Zfs.Juesequanxian(true)+zfzhi.Zhi.Xx()+zfzhi.Zhi.Shuzi1w(),
nil,

)
w:=httptest.NewRecorder()
beego.BeeApp.Handlers.ServeHTTP(w,r)
log.Println(w.Body.String())

}
func TestJuesequanxiandelete(t *testing.T){
r,_:=http.NewRequest(
zf.Zfs.DELETE(false),
zfzhi.Zhi.Xx()+zf.Zfs.Juesequanxian(true)+zfzhi.Zhi.Xx()+zfzhi.Zhi.Shuzi1w(),
nil,

)
w:=httptest.NewRecorder()
beego.BeeApp.Handlers.ServeHTTP(w,r)
log.Println(w.Body.String())

}
func TestJuesequanxianpost(t *testing.T){
canshu:=zfzhi.Zhi.Postjuesequanxian()
r,_:=http.NewRequest(
zf.Zfs.POST(false),
zfzhi.Zhi.Xx()+zf.Zfs.Juesequanxian(true),
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
func TestJuesequanxianpatch(t *testing.T){
canshu:=zfzhi.Zhi.Patchjuesequanxian()
r,_:=http.NewRequest(
zf.Zfs.PATCH(false),
zfzhi.Zhi.Xx()+zf.Zfs.Juesequanxian(true),
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
