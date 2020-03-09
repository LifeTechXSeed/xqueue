package test

import (
	"strconv"
	"testing"
	"xqueue/entity"
	"xqueue/util"
	"xqueue/worker"

	"github.com/stretchr/testify/assert"
)

var cmdSample = "echo hello world"

func beforeTestDequeue(eIns *entity.Entity, jobId int, workerName string) error {
	err := eIns.Redis.HSet(util.JobInfoPrefix+strconv.Itoa(jobId), util.JobCmdKey, cmdSample)
	if err != nil {
		return err
	}

	err = eIns.Redis.ZJobToQueue(util.JobQueueKey+":"+workerName, jobId, 1)
	if err != nil {
		return err
	}

	return nil
}

func afterTestDequeue(eIns *entity.Entity, jobId int) {
	_ = eIns.Redis.Del(util.JobInfoPrefix + strconv.Itoa(jobId))
	_ = eIns.Redis.Del(util.JobQueueKey)
}

func workerGetJob(agent *worker.WorkerAgent) <-chan string {
	cmd := agent.GetJob()
	c := make(chan string)
	c <- cmd

	return c
}

func TestDequeue(t *testing.T) {
	assert := assert.New(t)

	eIns := entity.CreateNewEntity()
	defer eIns.Release()

	jobSampleId := 1
	workerName := "worker1"
	err := beforeTestDequeue(eIns, jobSampleId, workerName)
	defer afterTestDequeue(eIns, jobSampleId)

	if err != nil {
		t.Fatal(err)
	}

	wAgent := worker.NewAgent(workerName, *eIns)
	cmd := wAgent.GetJob()

	queueLen, _ := eIns.Redis.ZCount(util.JobQueueKey, "+inf", "-inf")

	assert.Equal(cmd, cmdSample)
	assert.Equal(int(queueLen), 0)
}
