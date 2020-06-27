package server

import (
	"potatoengine/src/space"
)

type IServer interface {
	RegisterSpace(sp space.ISpace)
	Stop()
	Run()
	SpaceRun() bool
}
