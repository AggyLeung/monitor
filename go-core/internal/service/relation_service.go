package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/example/go-core/internal/model"
)

type RelationRepo interface {
	Create(ctx context.Context, rel *model.Relation) error
	List(ctx context.Context) ([]model.Relation, error)
	ListByCI(ctx context.Context, ciID string) ([]model.Relation, error)
}

type RelationService struct {
	repo RelationRepo
}

func NewRelationService(repo RelationRepo) *RelationService {
	return &RelationService{repo: repo}
}

func (s *RelationService) List(ctx context.Context) ([]model.Relation, error) {
	return s.repo.List(ctx)
}

func (s *RelationService) ListByCI(ctx context.Context, ciID string) ([]model.Relation, error) {
	return s.repo.ListByCI(ctx, ciID)
}

func (s *RelationService) Create(ctx context.Context, rel *model.Relation) error {
	if rel.ID == uuid.Nil {
		rel.ID = uuid.New()
	}
	return s.repo.Create(ctx, rel)
}
