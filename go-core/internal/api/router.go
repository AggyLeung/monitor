package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/example/go-core/internal/api/handler"
	"github.com/example/go-core/internal/api/middleware"
)

type Handlers struct {
	CI       *handler.CIHandler
	Relation *handler.RelationHandler
	Task     *handler.TaskHandler
	Auth     *handler.AuthHandler
	GraphSync *handler.GraphSyncHandler
}

func NewRouter(h Handlers, secret string, enforcer *casbin.Enforcer, db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.RequestLog(), middleware.RateLimit(200, 400))

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/login", h.Auth.Login)
	}

	protected := r.Group("/api/v1")
	protected.Use(middleware.JWT(secret), middleware.Authorize(enforcer), middleware.AuditMiddleware(db))
	{
		protected.GET("/cis", h.CI.List)
		protected.POST("/cis", h.CI.Create)
		protected.PUT("/cis/:id", h.CI.Update)
		protected.DELETE("/cis/:id", h.CI.Delete)

		protected.GET("/relations", h.Relation.List)
		protected.POST("/relations", h.Relation.Create)
		protected.GET("/topology/:id", h.Relation.GetTopology)
		protected.GET("/topology/:id/impact", h.Relation.ImpactAnalysis)

		protected.POST("/tasks/discovery", h.Task.PublishScan)
		protected.POST("/sync/callback", h.Task.SyncCallback)
		protected.GET("/sync/failed", h.GraphSync.ListFailed)
		protected.POST("/sync/failed/:id/retry", h.GraphSync.Retry)
	}
	return r
}
