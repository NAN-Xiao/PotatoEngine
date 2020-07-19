package account

import (
	"potatoengine/src/entity"
)

type Account struct {
	entity.Entity
	Longin bool
}

func TestName(t *Account) {
	t.Read()
}
//开始接受发送线程
