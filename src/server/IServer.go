package server

import (
	"potatoengine/src/space"
)

type IServer interface {
	RegisterSpace(sp space.ISpace)
	//Initialize()
	//Begin()
	Stop()
	Run()
	//RunSpace()
}
