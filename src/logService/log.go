package logService

import (
	"fmt"
	"potatoengine/src/globleTimer"
)

var errorLog chan  string
var	defaultLog chan string

func init()  {
	errorLog=make(chan string,1024)
	defaultLog=make(chan string,1024)
	globleTimer.RegiestTick(Tick)
}

func Log(ctx string)  {
	s:=fmt.Sprintf("%s \n",ctx)
	errorLog<-s
	defaultLog<-s
}
func LogError(ctx string)  {
	s:=fmt.Sprintf("%s \n",ctx)
	errorLog<-s
}
func Tick()  {
	if errorLog!=nil&& len(errorLog)>0{
		err:=<-errorLog
		fmt.Sprintf("Error::%s.\n",err)
	}
	if defaultLog!=nil&& len(defaultLog)>0{
		log:=<-defaultLog
		fmt.Sprintf("Log::%s.\n",log)
	}
}