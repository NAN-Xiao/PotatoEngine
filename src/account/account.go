package account

import (
	"potatoengine/src/client"
	"potatoengine/src/entity"
)

type Account struct {
	entity.Entity
	Longin bool
}


func NewAccount(cl *client.Client) *Account {
	ac := new (Account)
	ac.Longin=false
	ac.CreatEntity(cl.Conn_id)
	return ac
}
