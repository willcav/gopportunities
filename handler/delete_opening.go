package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willcav/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id: %s not found", id))
		return
	}

	// Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		errorMessage := "Error deleting opening"
		logger.Errorf("%s: %v", errorMessage, err.Error())
		sendError(ctx, http.StatusInternalServerError, errorMessage)
		return
	}

	sendSuccess(ctx, "Delete Opening", opening)
}
