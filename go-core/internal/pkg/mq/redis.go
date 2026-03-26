package mq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Task struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Payload   map[string]interface{} `json:"payload"`
	CreatedAt time.Time              `json:"created_at"`
}

type TaskPublisher struct {
	client *redis.Client
	stream string
}

func NewTaskPublisher(addr, password, stream string, db int) *TaskPublisher {
	return &TaskPublisher{
		client: redis.NewClient(&redis.Options{Addr: addr, Password: password, DB: db}),
		stream: stream,
	}
}

func (p *TaskPublisher) Publish(task *Task) (string, error) {
	ctx := context.Background()
	data, err := json.Marshal(task)
	if err != nil {
		return "", err
	}
	return p.client.XAdd(ctx, &redis.XAddArgs{
		Stream: p.stream,
		Values: map[string]interface{}{"data": data},
	}).Result()
}
