package agent

import (
	"potatoengine/src/client"
	"potatoengine/src/entity"
)

type Agent struct {
	entity.Entity
}



func NewAgent(cl *client.Client) *Agent {
	ag := &Agent{

	}
	return ag
}
