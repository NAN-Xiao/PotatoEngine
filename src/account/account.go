package account

import (
	"fmt"
	"potatoengine/src/entity"
)

type Account struct {
	entity.Entity
	Longin bool
}

func (this *Account)GetEntity() (entity.Entity,error){
	if &this.Entity!=nil{
		return  this.Entity,nil
	}
	return nil,fmt.Errorf("account not hava any entity")
}
//开始接受发送线程
