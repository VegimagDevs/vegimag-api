package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateSessionJSON struct {
	Email    string `json:"email" binding:"required_without=username"`
	Username string `json:"username" binding:"required_without=email"`
	Password string `json:"password" binding:"required"`
}

// @Summary Create a session
// @Accept json
// @Produce json
// @Router /sessions [post]
func (handlers *Handlers) CreateSession(ctx *gin.Context) {
	createSessionJSON := new(CreateSessionJSON)

	if err := ctx.ShouldBindJSON(createSessionJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  "coucou",
		"refreshToken": "tu veux voir ma bite ?",
	})
}
