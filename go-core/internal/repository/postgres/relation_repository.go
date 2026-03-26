package postgres

import (
	"context"

	"gorm.io/gorm"
	"github.com/example/go-core/internal/model"
)

type RelationRepository struct {
	db *gorm.DB
}

func NewRelationRepository(db *gorm.DB) *RelationRepository {
	return &RelationRepository{db: db}
}

func (r *RelationRepository) Create(ctx context.Context, rel *model.Relation) error {
	return r.db.WithContext(ctx).Create(rel).Error
}

func (r *RelationRepository) List(ctx context.Context) ([]model.Relation, error) {
	var out []model.Relation
	err := r.db.WithContext(ctx).Find(&out).Error
	return out, err
}

func (r *RelationRepository) ListByCI(ctx context.Context, ciID string) ([]model.Relation, error) {
	var out []model.Relation
	err := r.db.WithContext(ctx).
		Where("source_ci_id::text = ? OR target_ci_id::text = ?", ciID, ciID).
		Find(&out).Error
	return out, err
}
