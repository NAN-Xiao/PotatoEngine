package account

import (
	"potatoengine/src/client"
	"potatoengine/src/entity"
)

type Account struct {
	Entity entity.Entity
	Longin bool
}

func NewAccount(cl *client.Client) *Account {
	
	ac := &Account{
		Entity: entity.Entity{},
		Longin: false,
	}
	return ac
}