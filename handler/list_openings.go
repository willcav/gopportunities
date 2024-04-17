package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willcav/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		errorMessage := "Error fetching openings"
		logger.Errorf("%s: %v", errorMessage, err.Error())
		sendError(ctx, http.StatusInternalServerError, errorMessage)
		return
	}

	sendSuccess(ctx, "List Openings", openings)
}
