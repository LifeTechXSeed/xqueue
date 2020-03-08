package arbiter

import (
	"strconv"
	"xqueue/entity"
	"xqueue/log"
	"xqueue/util"
)

type Job struct {
	jid int
	command string
	priority int
}

var logger = log.NewLogger("worker-agent")

func enqueue(eIns entity.Entity, job Job) (int, error) {
	err := eIns.Redis.ZJobToQueue(util.JobQueueKey, job.jid, job.priority)
	if err != nil {
		logger.Error("error when add job to queue", err)
		return 0, err
	}

	err = eIns.Redis.HSet(util.JobInfoPrefix+strconv.Itoa(job.jid), util.JobCmdKey, job.command)
	if err != nil {
		logger.Error("error when add job command")
		return 0, err
	}

	return job.jid, nil
}
