package tests
import(
"github.com/astaxie/beego"

_"mhsydata/luyous"
"net/http"
."github.com/smartystreets/goconvey/convey"
"testing"
"mhsydata/chushihuas"
"net/http/httptest"
"changliang/zf"
"changliang/zfzhi"

)
func init(){
beego.TestBeegoInit(chushihuas.Getapppath())
}
func TestMainluyou(t *testing.T){
r,_:=http.NewRequest(
zf.Zfs.GET(false),
zfzhi.Zhi.Xx(),
nil,

)
w:=httptest.NewRecorder()
beego.BeeApp.Handlers.ServeHTTP(w,r)
beego.Trace(
zf.Zfs.Testing(true)+zfzhi.Zhi.Zkhz()+zfzhi.Zhi.Bfh()+zf.Zfs.D(true)+zfzhi.Zhi.Zkhy(),
zf.Zfs.TestMainluyou(false)+zfzhi.Zhi.Zkhz()+zfzhi.Zhi.Bfh()+zf.Zfs.D(true)+zfzhi.Zhi.Zkhy(),
zf.Zfs.Code(false)+zfzhi.Zhi.Zkhz()+zfzhi.Zhi.Bfh()+zf.Zfs.D(true)+zfzhi.Zhi.Zkhy()+zfzhi.Zhi.Fxx()+zfzhi.Zhi.Bfh()+zf.Zfs.S(true),
w.Code,
w.Code,
w.Code,
w.Body.String(),

)
Convey(
zf.Zfs.Code(false)+zfzhi.Zhi.Kgf()+zf.Zfs.And(false)+zf.Zfs.Bodylen(false),
t,
func(){
Convey(
zf.Zfs.Code(false)+zfzhi.Zhi.Dyh()+zfzhi.Zhi.Dyh()+zfzhi.Zhi.Shuzi2w()+zfzhi.Zhi.Shuzi0w()+zfzhi.Zhi.Shuzi0w(),
func(){
So(
w.Code,
ShouldEqual,
zfzhi.Zhi.Shuzi2()*zfzhi.Zhi.Shuzi10()*zfzhi.Zhi.Shuzi10(),

)

},

)
Convey(
zf.Zfs.Bodylen(false)+zfzhi.Zhi.Dy()+zfzhi.Zhi.Shuzi0w(),
func(){
So(
w.Body.Len(),
ShouldBeGreaterThan,
zfzhi.Zhi.Shuzi0(),

)

},

)

},

)

}
