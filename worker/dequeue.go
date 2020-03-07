package worker

import (
	"strconv"
	"xqueue/entity"
	"xqueue/util"
)

func dequeue(eIns entity.Entity) string {
	ids, err := eIns.Redis.ZRange(util.JobQueueKey, 0, 1)
	if err != nil {
		logger.Error("error when pop from queue ", err)
		return ""
	}

	if len(ids) <= 0 {
		return ""
	}

	id := ids[0]
	logger.Info("job id ", id)

	isLock, err := eIns.Redis.HGet(util.JobInfoPrefix+id, util.JobLockStateKey)
	if err != nil {
		logger.Error("error when check job lock state info")
		return ""
	}

	logger.Info("lock state: ", isLock)
	locked, _ := strconv.Atoi(isLock)
	if locked == 1 {
		logger.Info("job is locked, it's will run in other worker")
		return ""
	}

	cmd, err := eIns.Redis.HGet(util.JobInfoPrefix+id, util.JobCmdKey)
	if err != nil {
		logger.Error("error when get job cmd")
		return ""
	}

	err = eIns.Redis.HSet(util.JobInfoPrefix+id, util.JobLockStateKey, "1")
	if err != nil {
		logger.Error("error when lock job")
		return ""
	}

	err = eIns.Redis.ZRem(util.JobQueueKey, id)
	if err != nil {
		logger.Error("error when pop job from queue", err)
		return ""
	}

	return cmd
}
