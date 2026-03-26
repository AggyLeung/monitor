package service

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/example/go-core/internal/model"
)

type GraphSyncRepo interface {
	Enqueue(ctx context.Context, item *model.GraphSyncFailed) error
	ListByStatus(ctx context.Context, status string, limit int) ([]model.GraphSyncFailed, error)
	MarkRetried(ctx context.Context, id uuid.UUID, success bool, errorMessage string) error
}

type GraphSyncService struct {
	repo GraphSyncRepo
}

func NewGraphSyncService(repo GraphSyncRepo) *GraphSyncService {
	return &GraphSyncService{repo: repo}
}

func (s *GraphSyncService) EnqueueFailure(ctx context.Context, entityType string, entityID *uuid.UUID, payload interface{}, errMsg string) error {
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	item := &model.GraphSyncFailed{
		ID:           uuid.New(),
		EntityType:   entityType,
		EntityID:     entityID,
		Payload:      raw,
		ErrorMessage: errMsg,
		Status:       "pending",
	}
	return s.repo.Enqueue(ctx, item)
}

func (s *GraphSyncService) ListFailed(ctx context.Context, status string, limit int) ([]model.GraphSyncFailed, error) {
	return s.repo.ListByStatus(ctx, status, limit)
}

func (s *GraphSyncService) Retry(ctx context.Context, id uuid.UUID) error {
	// Placeholder retry execution: wire actual Neo4j upsert logic in worker/job.
	return s.repo.MarkRetried(ctx, id, true, "")
}
