package account

import "potatoengine/src/client"

type Account struct {
	_client *client.Client
}

func (this *Account) Recive() {

}

func NewAccount(cl *client.Client) *Account {
	if cl == nil {
		return nil
	}
	a := &Account{_client: cl}
	return a
}
