package test

import (
	"xqueue/log"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	log.LoggerInit()
}
