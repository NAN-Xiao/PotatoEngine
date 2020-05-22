package dispatcher

import "potatoengine/src/message"

type Dispatcher struct {
}

func DisposMessage(msg *message.Messsage) {
	if msg == nil {
		return
	}
	switch msg.GetID() {
	case uint32(message.Msg_Login):
		//todo
		//查找账号信息返回
		return
	case uint32(message.Msg_Regist):
		//todo
		//注册账号
		return
	case uint32(message.Msg_TEnter):
		//todo
		//成功进入游戏 发送角色信息
		return

	}
}
