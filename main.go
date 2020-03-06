package main

import (
	"time"
	"xqueue/entity"
	"xqueue/log"
	"xqueue/worker"
)

var logger = log.NewLogger("main")

func init() {
	log.LoggerInit()
}

func main() {
	logger.Info("hello")
	enIns := entity.CreateNewEntity()
	wAgent := worker.NewAgent("test", *enIns)
	defer wAgent.Stop()

	wAgent.Start()
	time.Sleep(time.Duration(1) * time.Minute)
}
