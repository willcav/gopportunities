package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willcav/gopportunities/schemas"
)

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
