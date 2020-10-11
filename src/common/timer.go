package common

import "time"

var tick *time.Ticker
var tickfunc []func()

//初始化
//30帧
func init() {
	//todo 后期从配置读取帧率
	tick = time.NewTicker(time.Second / 10)
	tickfunc = make([]func(), 0)
}

//注册需要tick的函数
func RegiestTick(f func()) {
	tickfunc = append(tickfunc, f)
}

//删除注册的tick
func UnRegistTick(f func()) {

	if len(tickfunc) <= 0 || tickfunc == nil {
		return
	}
	for i := range tickfunc {
		if &tickfunc[i]==&f{
			tickfunc = append(tickfunc[:i], tickfunc[i+1:]...)
		}
	}
}

//Tike计时
func Tick() {
	if tick == nil {
		return
	}
	go func() {
		//println("start tick")
		for {
			select {
			case <-tick.C:
				ln := len(tickfunc)
				if ln <= 0 {
					continue
				}
				for i := 0; i < ln; i++ {
					fn := tickfunc[i]
					fn()
				}
			}
		}
	}()
}
