package controllers

import (
	"booleans/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestBody is the body of the incoming request
type RequestBody struct {
	Key   string `json:"key"`
	Value *bool  `json:"value" binding:"required"`
}

// Controller stuff
type Controller interface {
	NewBoolean(ctx *gin.Context)
	GetBoolean(ctx *gin.Context)
	UpdateBoolean(ctx *gin.Context)
	DeleteBoolean(ctx *gin.Context)
}

// ControllerImpl is best
type ControllerImpl struct{}

// MyController rocks
var MyController = ControllerImpl{}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// NewBoolean will manage the creation of a new Boolean
func (ctrl ControllerImpl) NewBoolean(ctx *gin.Context) {
	var r RequestBody
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := services.MyRepo.AddToDB(r.Key, *r.Value)
	ctx.JSON(http.StatusOK, gin.H{
		"key":   result.Name,
		"value": result.Val,
		"id":    result.UUID,
	})
}

// GetBoolean will fetch an entry based upon the uuid provided
func (ctrl ControllerImpl) GetBoolean(ctx *gin.Context) {
	uuid := ctx.Param("id")
	if isValidUUID(uuid) == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The parameter is not a UUID"})
		return
	}
	result, err := services.MyRepo.GetFromDB(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"key":   result.Name,
			"value": result.Val,
			"id":    result.UUID,
		})
	}
}

// UpdateBoolean will make specified changes to the entry
func (ctrl ControllerImpl) UpdateBoolean(ctx *gin.Context) {
	uuid := ctx.Param("id")
	if isValidUUID(uuid) == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The parameter is not a UUID"})
		return
	}
	var r RequestBody
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := services.MyRepo.UpdateInDB(r.Key, *r.Value, uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"key":   result.Name,
			"value": result.Val,
			"id":    result.UUID,
		})
	}
}

// DeleteBoolean will delete the entry
func (ctrl ControllerImpl) DeleteBoolean(ctx *gin.Context) {
	uuid := ctx.Param("id")
	if isValidUUID(uuid) == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The parameter is not a UUID"})
		return
	}
	err := services.MyRepo.DeleteFromDB(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Record Not Found")
	} else {
		ctx.Writer.WriteHeader(http.StatusNoContent)
	}
}
