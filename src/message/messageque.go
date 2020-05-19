package message

type MessageQue struct {
	_que []*messsage
}

func (queue *MessageQue) Pop() *messsage {
	if len(queue._que) <= 0 {
		return nil
	}
	msg := queue._que[0]
	return msg
}

func (queue *MessageQue) PushBack(msg *messsage) {
	if msg == nil || queue._que == nil {
		return
	}
	queue._que = append(queue._que, msg)
}

func NewMessageQueue() *MessageQue {
	que := &MessageQue{_que: nil}
	return que
}
