package main

import (
	"xqueue/entity"
	"xqueue/log"
)

var logger = log.NewLogger("main")

func init() {
	log.LoggerInit()
}

func main() {
	logger.Info("hello")
	enIns := entity.CreateNewEntity()
	defer enIns.Release()

	// wAgent := worker.NewAgent("test", *enIns)
	// defer wAgent.Stop()

	// wAgent.Start()
	// time.Sleep(time.Duration(1) * time.Minute)
}
