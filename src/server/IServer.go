package server

import (
	"potatoengine/src/space"
)

type IServer interface {
	RegisterSpace(sp *space.BaseSpace)
	Initialize()
	Begin()
	Stop()
	Start()
}
