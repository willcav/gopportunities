package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willcav/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	// This bind populates the struct with the request body
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		errorMessage := "Error creating opening"
		logger.Errorf("%s: %v", errorMessage, err.Error())
		sendError(ctx, http.StatusInternalServerError, errorMessage)
		return
	}

	sendSuccess(ctx, "Create Opening", opening)
}
