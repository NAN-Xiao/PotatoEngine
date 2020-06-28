package dispatcher

import (
	"potatoengine/src/netmessage"
)

var Dispatch map[int32]func(pk *netmessage.MsgPackage)

//根据id处理消息

func init() {
	RegistDispathMap(0, nil)
}

//消除分发函数注册到map
func RegistDispathMap(id int32, f func(msgPackage *netmessage.MsgPackage)) {
	if Dispatch == nil {
		Dispatch = make(map[int32]func(msgPackage *netmessage.MsgPackage))
	}
	for _, ok := Dispatch[id]; ok == true; {
		return
	}
	Dispatch[id] = f

}

//从map得到对应消息处理函数并打包message进行处理
func DispatcherMessage(uid int32,pid int32,msg *netmessage.MsgPackage) {
	if Dispatch == nil {
		return
	}
	id,_:= netmessage.GetServerMsgID(msg)
	if _, ok := Dispatch[id]; ok == false {

		return
	}
	pkg := netmessage.PackMessagePackage(uid, pid, *msg)
	fc := Dispatch[id]
	fc(pkg)
}
