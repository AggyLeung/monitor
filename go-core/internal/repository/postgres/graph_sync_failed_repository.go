package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/example/go-core/internal/model"
)

type GraphSyncFailedRepository struct {
	db *gorm.DB
}

func NewGraphSyncFailedRepository(db *gorm.DB) *GraphSyncFailedRepository {
	return &GraphSyncFailedRepository{db: db}
}

func (r *GraphSyncFailedRepository) Enqueue(ctx context.Context, item *model.GraphSyncFailed) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *GraphSyncFailedRepository) ListByStatus(ctx context.Context, status string, limit int) ([]model.GraphSyncFailed, error) {
	var out []model.GraphSyncFailed
	q := r.db.WithContext(ctx).Order("created_at desc")
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if limit > 0 {
		q = q.Limit(limit)
	}
	err := q.Find(&out).Error
	return out, err
}

func (r *GraphSyncFailedRepository) MarkRetried(ctx context.Context, id uuid.UUID, success bool, errorMessage string) error {
	update := map[string]interface{}{
		"retry_count":    gorm.Expr("retry_count + 1"),
		"error_message":  errorMessage,
		"updated_at":     time.Now(),
		"next_retry_at":  nil,
	}
	if success {
		update["status"] = "resolved"
	} else {
		update["status"] = "pending"
		next := time.Now().Add(5 * time.Minute)
		update["next_retry_at"] = next
	}
	return r.db.WithContext(ctx).Model(&model.GraphSyncFailed{}).Where("id = ?", id).Updates(update).Error
}
