package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/example/go-core/internal/model"
)

type CIRepo interface {
	List(ctx context.Context) ([]model.CI, error)
	Get(ctx context.Context, id uuid.UUID) (*model.CI, error)
	Create(ctx context.Context, ci *model.CI) error
	Update(ctx context.Context, ci *model.CI) error
	SoftDelete(ctx context.Context, id uuid.UUID) error
}

type CIService struct {
	repo CIRepo
}

func NewCIService(repo CIRepo) *CIService {
	return &CIService{repo: repo}
}

func (s *CIService) List(ctx context.Context) ([]model.CI, error) {
	return s.repo.List(ctx)
}

func (s *CIService) Create(ctx context.Context, ci *model.CI) error {
	if ci.ID == uuid.Nil {
		ci.ID = uuid.New()
	}
	if ci.Status == "" {
		ci.Status = "active"
	}
	return s.repo.Create(ctx, ci)
}

func (s *CIService) Update(ctx context.Context, ci *model.CI) error {
	return s.repo.Update(ctx, ci)
}

func (s *CIService) SoftDelete(ctx context.Context, id uuid.UUID) error {
	return s.repo.SoftDelete(ctx, id)
}
