package dispatcher

import (
	"potatoengine/src/netmessage"
)

type Dispatcher struct {
}

var Dispatch map[int32]func(pk *netmessage.ServerMsgPackage)

//从map得到对应消息处理函数并打包message进行处理
func (this *Dispatcher) DispatcherMessage(uid int32, pid int32, msg *netmessage.MsgPackage) {
	if Dispatch == nil {
		return
	}
	id, _ := netmessage.GetServerMsgID(msg)
	if _, ok := Dispatch[id]; ok == false {

		return
	}
	pkg := netmessage.PackMessagePackage(uid, pid, *msg)
	fc := Dispatch[id]
	fc(pkg)
}
