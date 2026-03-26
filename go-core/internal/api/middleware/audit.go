package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/example/go-core/internal/model"
)

func AuditMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Request.Method == "GET" {
			return
		}

		var affected *uuid.UUID
		if ciID, ok := c.Get("affected_ci_id"); ok {
			switch v := ciID.(type) {
			case uuid.UUID:
				affected = &v
			case string:
				if parsed, err := uuid.Parse(v); err == nil {
					affected = &parsed
				}
			}
		}

		action := "update"
		switch c.Request.Method {
		case "POST":
			action = "create"
		case "PUT", "PATCH":
			action = "update"
		case "DELETE":
			action = "delete"
		}

		log := model.AuditLog{
			ID:        uuid.New(),
			Action:    action,
			CIID:      affected,
			OldValue:  []byte("null"),
			NewValue:  []byte("null"),
			CreatedAt: time.Now(),
		}
		go func(entry model.AuditLog) {
			_ = db.Create(&entry).Error
		}(log)
	}
}
