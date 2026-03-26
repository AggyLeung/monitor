package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/example/go-core/internal/model"
	"github.com/example/go-core/internal/service"
)

type RelationHandler struct {
	relSvc *service.RelationService
	topo   *service.TopologyService
}

func NewRelationHandler(relSvc *service.RelationService, topo *service.TopologyService) *RelationHandler {
	return &RelationHandler{relSvc: relSvc, topo: topo}
}

func (h *RelationHandler) List(c *gin.Context) {
	items, err := h.relSvc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *RelationHandler) Create(c *gin.Context) {
	var req model.Relation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.relSvc.Create(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *RelationHandler) GetTopology(c *gin.Context) {
	ciID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	depth := 2
	if q := c.Query("depth"); q != "" {
		if d, e := strconv.Atoi(q); e == nil {
			depth = d
		}
	}
	nodes, edges, err := h.topo.GetTopology(c.Request.Context(), ciID, depth)
	if err != nil {
		// Fallback: use PostgreSQL relation table when Neo4j is unavailable.
		rels, relErr := h.relSvc.ListByCI(c.Request.Context(), ciID.String())
		if relErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "fallback_error": relErr.Error()})
			return
		}
		fallbackEdges := make([]service.Edge, 0, len(rels))
		for _, rel := range rels {
			fallbackEdges = append(fallbackEdges, service.Edge{
				Source: rel.SourceCIID.String(),
				Target: rel.TargetCIID.String(),
				Type:   rel.Type,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"nodes":        []service.Node{},
			"edges":        fallbackEdges,
			"degraded":     true,
			"degrade_from": "neo4j_to_postgres_relation",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"nodes": nodes, "edges": edges})
}

func (h *RelationHandler) ImpactAnalysis(c *gin.Context) {
	ciID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	depth := 5
	if q := c.Query("depth"); q != "" {
		if d, e := strconv.Atoi(q); e == nil {
			depth = d
		}
	}
	nodes, edges, err := h.topo.ImpactAnalysis(c.Request.Context(), ciID, depth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"nodes": nodes, "edges": edges, "query": "impact_analysis"})
}
