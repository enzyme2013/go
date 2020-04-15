package User

import (
	"common/event"
	"common/sql"
	"log"
	"protos"
	"slg/const"
	"slg/entity"
)

//event--------------------------------------
func init() {
	Event.Reg(Const.OnInitDB, func() {
		x := Sql.ORM()
		x.Sync2(new(Entity.User))
	})
	Event.Reg(Const.OnUserNew, func(uid int64) {
		log.Println("User.OnUserNew", uid)
	})
	Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {
		log.Println("User.OnUserGetData", uid)
	})
}
