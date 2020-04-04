package test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"xqueue/arbiter"
	"xqueue/entity"
	"xqueue/util"
)

func afterTestEnqueue(eIns *entity.Entity, jobId int) {
	eIns.Redis.Del(util.JobInfoPrefix + strconv.Itoa(jobId))
	eIns.Redis.Del(util.JobQueueKey)
}

func TestEnqueue(t *testing.T) {
	assert := assert.New(t)

	eIns := entity.CreateNewEntity()
	defer eIns.Release()

	jobSampleId := 1
	jobSample := arbiter.Job{Jid: jobSampleId, Command: "echo HELLO_WORLD", Priority: 1}
	jid, err := arbiter.Enqueue(*eIns, jobSample)
	if err != nil {
		t.Fatal(err)
	}
	defer afterTestEnqueue(eIns, jobSampleId)

	jobSampleId = 2
	jobSample = arbiter.Job{Jid: jobSampleId, Command: "echo HELLO_WORLD 2", Priority: 1}
	jid, err = arbiter.Enqueue(*eIns, jobSample)
	if err != nil {
		t.Fatal(err)
	}

	jobSampleId = 3
	jobSample = arbiter.Job{Jid: jobSampleId, Command: "echo HELLO_WORLD 3", Priority: 3}
	jid, err = arbiter.Enqueue(*eIns, jobSample)
	if err != nil {
		t.Fatal(err)
	}

	ids, err := eIns.Redis.ZRange(util.JobQueueKey, 0, -1)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"1", "2", "3"}

	assert.Equal(3, jid)
	assert.Equal(expected, ids)
}
