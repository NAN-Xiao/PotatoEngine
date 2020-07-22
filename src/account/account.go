package account

import (
	"potatoengine/src/entity"
)

type Account struct {
	entity.Entity
	entity.IEntity
	Longin bool
}

func (this *Account)GetEntity() entity.Entity {
	return this.Entity
}
//开始接受发送线程
