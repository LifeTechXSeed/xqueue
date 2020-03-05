package main

import (
	"xqueue/log"
)

var logger = log.NewLogger("main")

func init() {
	log.LoggerInit()
}

func main() {
	logger.Info("hello")
}
