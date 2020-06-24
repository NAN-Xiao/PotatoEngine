package message

var MessageMap map[int32]interface{}

func init() {
	MessageMap = make(map[int32]interface{})
}

//注册消息到messageMap
func RegisteMessageID(msg interface{}) {
	id, ok := GetServerMsgID(msg)
	if ok != nil {
		return
	}
	_, has := MessageMap[id]
	if has {
		return
	}
	MessageMap[id] = msg
}

