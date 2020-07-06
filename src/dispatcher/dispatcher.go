package dispatcher

import (
	"potatoengine/src/netmessage"
)

type Dispatcher struct {
}

var Dispatch map[int32]func(pk *netmessage.ServerMsgPackage)

//从map得到对应消息处理函数并打包message进行处理
func (this *Dispatcher) Dispatch(m *netmessage.ServerMsgPackage) {
	mid, err := netmessage.GetServerMsgID(m.Msg)
	if err != nil {
		return
	}
	fn := netmessage.GetProcessFuction(mid)
	fn(m)
}
