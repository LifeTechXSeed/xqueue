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

	//ids, err := eIns.Redis.ZRange(util.JobQueueKey, 0, -1)
	//if err != nil {
	//	logger.Error("bla ", err)
	//	return 0, err
	//}
	//fmt.Println(ids)
	//
	//queueLen, err := eIns.Redis.ZCount(util.JobQueueKey, "+inf", "-inf")
	//if err != nil {
	//	logger.Error("foo ", err)
	//	return 0, err
	//}
	//fmt.Println(queueLen)

	err = eIns.Redis.HSet(util.JobInfoPrefix+strconv.Itoa(job.Jid), util.JobCmdKey, job.Command)
	if err != nil {
		logger.Error("error when add job command")
		return 0, err
	}

	return job.Jid, nil
}
