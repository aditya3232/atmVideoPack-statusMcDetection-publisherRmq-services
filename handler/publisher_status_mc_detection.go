package handler

import (
	"net/http"

	"github.com/aditya3232/gatewatchApp-services.git/constant"
	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/log"
	"github.com/aditya3232/gatewatchApp-services.git/model/publisher_status_mc_detection"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PublisherStatusMcDetectionHandler struct {
	publisherStatusMcDetectionService publisher_status_mc_detection.Service
}

func NewPublisherStatusMcDetectionHandler(publisherStatusMcDetectionService publisher_status_mc_detection.Service) *PublisherStatusMcDetectionHandler {
	return &PublisherStatusMcDetectionHandler{publisherStatusMcDetectionService}
}

// public message to rmqg
func (h *PublisherStatusMcDetectionHandler) CreateQueueStatusMcDetection(c *gin.Context) {
	var input publisher_status_mc_detection.PublisherStatusMcInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	statusMcDetection, err := h.publisherStatusMcDetectionService.CreateQueueStatusMcDetection(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, publisher_status_mc_detection.FormatPublisherStatusMcDetection(statusMcDetection))
	log.Info("Queue status mc detection berhasil dibuat")
	c.JSON(http.StatusOK, response)
}
