package chushihuas
func init(){
Shezhilie()
Guojihualie()

}
var Chushihuas=make(map[string]Tongyong)
var Shujukus=make(map[string]Tongyong)
var Cuowus=make(map[string]Tongyong)
var Tishis=make(map[string]Tongyong)
func Shezhilie(){
shezhi:=Shezhijson()
for _,chushihua:=range shezhi.Chushihua{
Chushihuas[chushihua.Bianma]=chushihua
}
for _,shujuku:=range shezhi.Shujuku{
Shujukus[shujuku.Bianma]=shujuku
}

}
func Guojihualie(){
guojihua:=Guojihuajson()
for _,cuowu:=range guojihua.Cuowu{
Cuowus[cuowu.Bianma]=cuowu
}
for _,tishi:=range guojihua.Tishi{
Tishis[tishi.Bianma]=tishi
}

}
