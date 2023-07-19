package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sindhutrisha/sindhu/teams/pkg/rest/server/models"
	"github.com/sindhutrisha/sindhu/teams/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TrishaController struct {
	trishaService *services.TrishaService
}

func NewTrishaController() (*TrishaController, error) {
	trishaService, err := services.NewTrishaService()
	if err != nil {
		return nil, err
	}
	return &TrishaController{
		trishaService: trishaService,
	}, nil
}

func (trishaController *TrishaController) CreateTrisha(context *gin.Context) {
	// validate input
	var input models.Trisha
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger trisha creation
	if _, err := trishaController.trishaService.CreateTrisha(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Trisha created successfully"})
}

func (trishaController *TrishaController) UpdateTrisha(context *gin.Context) {
	// validate input
	var input models.Trisha
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger trisha update
	if _, err := trishaController.trishaService.UpdateTrisha(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Trisha updated successfully"})
}

func (trishaController *TrishaController) FetchTrisha(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger trisha fetching
	trisha, err := trishaController.trishaService.GetTrisha(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, trisha)
}

func (trishaController *TrishaController) DeleteTrisha(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger trisha deletion
	if err := trishaController.trishaService.DeleteTrisha(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Trisha deleted successfully",
	})
}

func (trishaController *TrishaController) ListTrishas(context *gin.Context) {
	// trigger all trishas fetching
	trishas, err := trishaController.trishaService.ListTrishas()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, trishas)
}

func (*TrishaController) PatchTrisha(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*TrishaController) OptionsTrisha(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*TrishaController) HeadTrisha(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
