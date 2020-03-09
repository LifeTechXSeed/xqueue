package worker

import (
	"xqueue/entity"
	"xqueue/util"
)

func dequeue(eIns entity.Entity, workerQueue string) string {
	ids, err := eIns.Redis.ZRange(workerQueue, 0, 1)
	if err != nil {
		logger.Error("error when pop from queue ", err)
		return ""
	}

	if len(ids) <= 0 {
		return ""
	}

	id := ids[0]
	logger.Info("job id ", id)

	cmd, err := eIns.Redis.HGet(util.JobInfoPrefix+id, util.JobCmdKey)
	if err != nil {
		logger.Error("error when get job cmd")
		return ""
	}

	err = eIns.Redis.ZRem(util.JobQueueKey, id)
	if err != nil {
		logger.Error("error when pop job from queue", err)
		return ""
	}

	return cmd
}
