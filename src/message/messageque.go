package message

type MessageQue struct {
	_que []*Messsage
}

//返回顶部msg并从队列移除
func (queue *MessageQue) Pop() *Messsage {
	if len(queue._que) <= 0 {
		return nil
	}
	msg := queue._que[0]
	queue._que = queue._que[1:]
	return msg
}

//消息放入队列后面
func (queue *MessageQue) PushBack(msg *Messsage) {
	if msg == nil {
		return
	}
	queue._que = append(queue._que, msg)
}

//创建一个新的消息队列，参数是新的消息数组长度
func NewMessageQueue(len uint32) *MessageQue {
	que := &MessageQue{_que: make([]*Messsage, len)}
	return que
}
