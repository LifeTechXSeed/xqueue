package worker

import "xqueue/entity"

type WorkerAgent struct {
	entityIns *entity.Entity
}

func NewAgent(eIns entity.Entity) *WorkerAgent {
	return &WorkerAgent{
		entityIns: &eIns,
	}
}
