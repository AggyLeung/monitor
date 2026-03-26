package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/example/go-core/internal/pkg/mq"
)

type TaskService struct {
	publisher *mq.TaskPublisher
}

func NewTaskService(pub *mq.TaskPublisher) *TaskService {
	return &TaskService{publisher: pub}
}

func (s *TaskService) PublishScanTask(scope string) (string, error) {
	task := &mq.Task{
		ID:        uuid.NewString(),
		Type:      "discovery.scan",
		CreatedAt: time.Now(),
		Payload: map[string]interface{}{
			"scope": scope,
		},
	}
	return s.publisher.Publish(task)
}
