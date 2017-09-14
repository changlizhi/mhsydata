package tests

import (
	"github.com/astaxie/beego/orm"

	"log"
	"mhsydata/chushihuas"
	"testing"
)

func TestDayin(t *testing.T) {
	log.Println("chushihuas.Defaultormer():====", chushihuas.Defaultormer())
	log.Println("chushihuas.Chushihuas:====", chushihuas.Chushihuas)
	log.Println("chushihuas.Shujukus:====", chushihuas.Shujukus)
	log.Println("chushihuas.Tishis:====", chushihuas.Tishis)
	log.Println("chushihuas.Cuowus:====", chushihuas.Cuowus)
	log.Println("orm.Debug:====", orm.Debug)
	log.Println("chushihuas.Getapppath():====", chushihuas.Getapppath())

}
