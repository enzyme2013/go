package Entity

import (
	"protos"

	"github.com/golang/protobuf/proto"
)

//角色数据
type User struct {
	Uid        int64  `xorm:"pk autoincr"`
	Passport   string `xorm:"index"`
	Name       string `xorm:"unique"`
	Gender     int32
	Head       int32
	HeadB      int32
	Level      int32 //`xorm:"index"`
	CityX      int32 `xorm:"index(pos)"`
	CityY      int32 `xorm:"index(pos)"`
	CreateTime int64 `xorm:"created"`
	LoginTime  int64 `xorm:"->"`
	AllianceId int64
	Version    int `xorm:"version"` //乐观锁
}

//转proto对象
func (this *User) ToProto() *protos.User {
	return &protos.User{
		Uid:        proto.Int64(this.Uid),
		Name:       proto.String(this.Name),
		Gender:     proto.Int32(this.Gender),
		Icon:       proto.Int32(this.Head),
		IconB:      proto.Int32(this.HeadB),
		Level:      proto.Int32(this.Level),
		CityX:      proto.Int32(this.CityX),
		CityY:      proto.Int32(this.CityY),
		AllianceId: proto.Int64(this.AllianceId),
	}
}

//转proto前端主键对象
func (this *User) ToProtoPK() *protos.UserPK {
	return &protos.UserPK{
		Uid: proto.Int64(this.Uid),
	}
}

//加到更新
func (this *User) AppendTo(updates *protos.Updates) {
	list := updates.User
	if list == nil {
		list = []*protos.User{}
	}
	updates.User = append(list, this.ToProto())
}

//加到删除
// func (this *User) AppendToPK(removes *protos.Removes) {
// 	list := removes.User
// 	if list == nil {
// 		list = []*protos.UserPK{}
// 	}
// 	removes.User = append(list, this.ToProtoPK())
// }
