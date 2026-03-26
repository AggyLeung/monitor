package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/example/go-core/internal/service"
)

type GraphSyncHandler struct {
	svc *service.GraphSyncService
}

func NewGraphSyncHandler(svc *service.GraphSyncService) *GraphSyncHandler {
	return &GraphSyncHandler{svc: svc}
}

func (h *GraphSyncHandler) ListFailed(c *gin.Context) {
	status := c.Query("status")
	limit := 50
	if q := c.Query("limit"); q != "" {
		if parsed, err := strconv.Atoi(q); err == nil && parsed > 0 && parsed <= 500 {
			limit = parsed
		}
	}
	items, err := h.svc.ListFailed(c.Request.Context(), status, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *GraphSyncHandler) Retry(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Retry(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"retried": true, "id": id})
}
