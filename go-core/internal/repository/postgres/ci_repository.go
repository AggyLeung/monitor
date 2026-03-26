package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/example/go-core/internal/model"
)

type CIRepository struct {
	db *gorm.DB
}

func NewCIRepository(db *gorm.DB) *CIRepository {
	return &CIRepository{db: db}
}

func (r *CIRepository) List(ctx context.Context) ([]model.CI, error) {
	var out []model.CI
	err := r.db.WithContext(ctx).Where("status <> ?", "deleted").Find(&out).Error
	return out, err
}

func (r *CIRepository) Get(ctx context.Context, id uuid.UUID) (*model.CI, error) {
	var ci model.CI
	if err := r.db.WithContext(ctx).First(&ci, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &ci, nil
}

func (r *CIRepository) Create(ctx context.Context, ci *model.CI) error {
	return r.db.WithContext(ctx).Create(ci).Error
}

func (r *CIRepository) Update(ctx context.Context, ci *model.CI) error {
	return r.db.WithContext(ctx).Model(&model.CI{}).Where("id = ?", ci.ID).Updates(ci).Error
}

func (r *CIRepository) SoftDelete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&model.CI{}).Where("id = ?", id).Update("status", "deleted").Error
}
