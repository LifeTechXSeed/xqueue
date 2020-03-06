package test

import (
	"testing"
	"time"
	"xqueue/entity"
	"xqueue/util"
	"xqueue/worker"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)
	entityIns := entity.CreateNewEntity()
	defer entityIns.Release()

	wAgent := worker.NewAgent("test", *entityIns)
	channel := util.GetEnv("ARBITER_CHANNEL", "xqueue")

	pubsub := entityIns.Redis.Subscribe(channel)

	_, err := pubsub.Receive()
	if err != nil {
		t.Fatal(err)
	}

	ch := pubsub.Channel()
	wAgent.HealthCheckReport()

	time.AfterFunc(time.Second, func() {
		_ = pubsub.Close()
	})

	for msg := range ch {
		data := worker.MessagePublish{}
		err = json.Unmarshal([]byte(msg.Payload), &data)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(channel, msg.Channel)
		assert.Equal(data.Name, "test")
		assert.Equal(data.MsgType, worker.MsgHealthCheck)
		assert.NotNil(data.Data)
	}
}
