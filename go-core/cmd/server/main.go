package main

import (
	"context"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/example/go-core/internal/api"
	"github.com/example/go-core/internal/api/handler"
	"github.com/example/go-core/internal/model"
	neo4jrepo "github.com/example/go-core/internal/repository/neo4j"
	"github.com/example/go-core/internal/repository/postgres"
	"github.com/example/go-core/internal/service"
	"github.com/example/go-core/internal/pkg/auth"
	"github.com/example/go-core/internal/pkg/mq"
)

func env(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func main() {
	pgDSN := env("POSTGRES_DSN", "host=127.0.0.1 user=postgres password=postgres dbname=cmdb port=5432 sslmode=disable")
	jwtSecret := env("JWT_SECRET", "replace-me")
	redisAddr := env("REDIS_ADDR", "127.0.0.1:6379")
	redisPwd := env("REDIS_PASSWORD", "")
	redisStream := env("REDIS_STREAM", "cmdb_tasks")
	neo4jURI := env("NEO4J_URI", "neo4j://127.0.0.1:7687")
	neo4jUser := env("NEO4J_USER", "neo4j")
	neo4jPwd := env("NEO4J_PASSWORD", "password")

	db, err := gorm.Open(postgres.Open(pgDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("postgres connect failed: %v", err)
	}
	if err := db.AutoMigrate(
		&model.CIType{},
		&model.CITypeAttribute{},
		&model.CI{},
		&model.CIAttributeValue{},
		&model.Relation{},
		&model.AuditLog{},
		&model.SyncTask{},
		&model.GraphSyncFailed{},
		&model.User{},
	); err != nil {
		log.Fatalf("postgres migrate failed: %v", err)
	}

	driver, err := neo4j.NewDriverWithContext(neo4jURI, neo4j.BasicAuth(neo4jUser, neo4jPwd, ""))
	if err != nil {
		log.Fatalf("neo4j connect failed: %v", err)
	}
	defer func() { _ = driver.Close(context.Background()) }()
	_ = neo4jrepo.New(driver)

	enforcer := auth.InitCasbin(db)
	ciSvc := service.NewCIService(postgres.NewCIRepository(db))
	relSvc := service.NewRelationService(postgres.NewRelationRepository(db))
	topoSvc := service.NewTopologyService(driver)
	taskSvc := service.NewTaskService(mq.NewTaskPublisher(redisAddr, redisPwd, redisStream, 0))
	graphSyncSvc := service.NewGraphSyncService(postgres.NewGraphSyncFailedRepository(db))

	h := api.Handlers{
		CI:       handler.NewCIHandler(ciSvc),
		Relation: handler.NewRelationHandler(relSvc, topoSvc),
		Task:     handler.NewTaskHandler(taskSvc),
		Auth:     handler.NewAuthHandler(jwtSecret),
		GraphSync: handler.NewGraphSyncHandler(graphSyncSvc),
	}
	r := api.NewRouter(h, jwtSecret, enforcer, db)

	addr := env("HTTP_ADDR", ":8080")
	log.Printf("go-core listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
