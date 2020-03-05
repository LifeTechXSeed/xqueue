package log

import (
	"io/ioutil"

	"xqueue/util"

	log "github.com/sirupsen/logrus"
)

var serviceLog = log.WithFields(log.Fields{"service": "service"})

type Debug struct {
	*log.Entry
}

func LoggerInit() {
	env := util.GetEnv("ENV", "Production")

	if env == "Production" {
		log.SetFormatter(&log.JSONFormatter{
			DisableTimestamp: false,
			TimestampFormat:  "2006-01-02 15:04:05",
		})

		return
	}

	if env == "Debug" {
		log.SetReportCaller(true)
	}
	if env == "Test" {
		log.SetOutput(ioutil.Discard)
		return
	}

	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       false,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
	})
}

func NewLogger(method string) *Debug {
	logger := serviceLog.WithFields(log.Fields{"method": method})
	return &Debug{logger}
}
