package main

import (
	"github.com/ijidan/jproto/jproto/call"
	log "github.com/sirupsen/logrus"
)

//入口函数
func main()  {
	go call.StartCallServer()
	log.Info("call server running...")
	//main阻塞
	ch:=make(chan struct{})
	<-ch
}
