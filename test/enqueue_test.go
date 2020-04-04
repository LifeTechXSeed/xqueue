package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"xqueue/arbiter"
	"xqueue/entity"
	"xqueue/util"
)

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

	queueLen, err := eIns.Redis.ZCount(util.JobQueueKey, "+inf", "-inf")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(3, jid)
	assert.Equal(3, int(queueLen))
}
