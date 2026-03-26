package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/example/go-core/internal/pkg/auth"
)

type AuthHandler struct {
	secret string
}

func NewAuthHandler(secret string) *AuthHandler {
	return &AuthHandler{secret: secret}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Placeholder auth: wire to user repository in next stage.
	role := "viewer"
	if req.Username == "admin" {
		role = "admin"
	}
	token, err := auth.GenerateToken(h.secret, req.Username, role, 12*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "role": role})
}
