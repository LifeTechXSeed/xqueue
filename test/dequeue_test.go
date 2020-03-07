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

func beforeTestDequeue(eIns *entity.Entity, jobId int) error {
	err := eIns.Redis.HSet(util.JobInfoPrefix+strconv.Itoa(jobId), util.JobCmdKey, cmdSample)
	if err != nil {
		return err
	}

	err = eIns.Redis.ZJobToQueue(util.JobQueueKey, jobId, 1)
	if err != nil {
		return err
	}

	return nil
}

func afterTestDequeue(eIns *entity.Entity, jobId int) {
	eIns.Redis.Del(util.JobInfoPrefix + strconv.Itoa(jobId))
	eIns.Redis.Del(util.JobQueueKey)
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
	err := beforeTestDequeue(eIns, jobSampleId)
	defer afterTestDequeue(eIns, jobSampleId)

	if err != nil {
		t.Fatal(err)
	}

	wAgent := worker.NewAgent("worker1", *eIns)
	cmd := wAgent.GetJob()

	lockState, _ := eIns.Redis.HGet(util.JobInfoPrefix+strconv.Itoa(jobSampleId), util.JobLockStateKey)
	queueLen, _ := eIns.Redis.ZCount(util.JobQueueKey, "+inf", "-inf")

	assert.Equal(cmd, cmdSample)
	assert.Equal(lockState, "1")
	assert.Equal(int(queueLen), 0)
}

// func TestDequeueParallel(t *testing.T) {
// 	assert := assert.New(t)

// 	eIns := entity.CreateNewEntity()
// 	defer eIns.Release()

// 	jobSampleId := 2
// 	err := beforeTestDequeue(eIns, jobSampleId)
// 	defer afterTestDequeue(eIns, jobSampleId)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	wAgent1 := worker.NewAgent("worker1", *eIns)
// 	wAgent2 := worker.NewAgent("worker2", *eIns)

// 	chanCmd1, chanCmd2 := workerGetJob(wAgent1), workerGetJob(wAgent2)
// 	cmd1, cmd2 := <-chanCmd1, <-chanCmd2

// 	cmdAssert := (cmd1 == "" && cmd2 == cmdSample) || (cmd1 == cmdSample && cmd2 == "")
// 	assert.Equal(cmdAssert, true)
// }
