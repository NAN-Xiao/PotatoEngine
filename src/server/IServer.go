package server



type IServer interface {
	Initialize()
	Begin()
	Stop()
	Serve()
}
