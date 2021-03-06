package worker

import (
	"xqueue/entity"
	"xqueue/log"
	"xqueue/util"

	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
)

const (
	MsgHealthCheck = 1
)

var logger = log.NewLogger("worker-agent")
var json = jsoniter.ConfigCompatibleWithStandardLibrary
var arbiterChannel = util.GetEnv("ARBITER_CHANNEL", "xqueue")

type WorkerAgent struct {
	name        string
	entityIns   *entity.Entity
	channel     *redis.PubSub
	intervalRun []chan bool
	queueKey    string
}

type MessagePublish struct {
	Name    string      `json:"worker_name"`
	MsgType int         `json:"msg_type"`
	Data    interface{} `json:"data"`
}

func NewAgent(name string, eIns entity.Entity) *WorkerAgent {
	channel := eIns.Redis.Subscribe(arbiterChannel)
	return &WorkerAgent{
		name:        name,
		entityIns:   &eIns,
		channel:     channel,
		intervalRun: []chan bool{},
		queueKey:    util.JobQueueKey + ":" + name,
	}
}

func (a *WorkerAgent) Start() {
	healtcheckRun := util.SetInterval(
		a.HealthCheckReport,
		5000,
		false,
	)
	a.intervalRun = append(a.intervalRun, healtcheckRun)
}

func (a *WorkerAgent) HealthCheckReport() {
	info := HealthCheck()
	rawMsg := MessagePublish{
		Name:    a.name,
		MsgType: MsgHealthCheck,
		Data:    info,
	}
	msg, err := json.Marshal(rawMsg)
	if err != nil {
		logger.Error("error when stringify report")
		return
	}
	err = a.entityIns.Redis.Publish(arbiterChannel, string(msg))
	if err != nil {
		logger.Error("error when publish message ")
	}
}

func (a *WorkerAgent) GetJob() string {
	return dequeue(*a.entityIns, a.queueKey)
}

func (a *WorkerAgent) Stop() {
	for _, run := range a.intervalRun {
		run <- true
	}
	_ = a.entityIns.Redis.Del(a.queueKey)
	a.channel.Close()
}
