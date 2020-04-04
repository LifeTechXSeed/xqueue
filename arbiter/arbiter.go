package arbiter

import (
	"strconv"
	"xqueue/entity"
	"xqueue/log"
	"xqueue/util"
)

type Job struct {
	Jid int
	Command string
	Priority int
}

var logger = log.NewLogger("worker-agent")

func Enqueue(eIns entity.Entity, job Job) (int, error) {
	err := eIns.Redis.ZJobToQueue(util.JobQueueKey, job.Jid, job.Priority)
	if err != nil {
		logger.Error("error when add job to queue", err)
		return 0, err
	}

	err = eIns.Redis.HSet(util.JobInfoPrefix+strconv.Itoa(job.Jid), util.JobCmdKey, job.Command)
	if err != nil {
		logger.Error("error when add job command")
		err := eIns.Redis.ZRem(util.JobQueueKey, strconv.Itoa(job.Jid))
		if err != nil {
			logger.Error("error when remove job queue")
		}
		return 0, err
	}

	return job.Jid, nil
}
