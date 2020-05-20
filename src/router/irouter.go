package router

type IRouter interface {
	PreDispose()
	Dispose()
	PostDisPose()
}
