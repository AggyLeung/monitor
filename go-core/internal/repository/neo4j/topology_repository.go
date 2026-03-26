package neo4jrepo

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repository struct {
	driver neo4j.DriverWithContext
}

func New(driver neo4j.DriverWithContext) *Repository {
	return &Repository{driver: driver}
}

func (r *Repository) Driver() neo4j.DriverWithContext {
	return r.driver
}

func (r *Repository) Ping(ctx context.Context) error {
	return r.driver.VerifyConnectivity(ctx)
}
