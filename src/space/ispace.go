package space

type ISpace interface {
	GetSpace() *BaseSpace
	//space启动的时候调用
	OnStart()
	//不按时间同步调用
	Process()
	//按时间间隔用
	Tick()
}
