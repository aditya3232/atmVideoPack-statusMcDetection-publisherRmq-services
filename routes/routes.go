package routes

import (
	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/aditya3232/gatewatchApp-services.git/connection"
	"github.com/aditya3232/gatewatchApp-services.git/handler"
	"github.com/aditya3232/gatewatchApp-services.git/middleware"
	"github.com/aditya3232/gatewatchApp-services.git/model/publisher_status_mc_detection"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	publisherStatusMcDetectionRepository := publisher_status_mc_detection.NewRepository(connection.RabbitMQ())

	// Initialize services
	publisherStatusMcDetectionService := publisher_status_mc_detection.NewService(publisherStatusMcDetectionRepository)

	// Initialize handlers
	publisherStatusMcDetectionHandler := handler.NewPublisherStatusMcDetectionHandler(publisherStatusMcDetectionService)

	// Configure routes
	api := router.Group("/publisher/atmvideopack/v1")

	statusMcDetectionRoutes := api.Group("/statusmcdetection", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))

	configureStatusMcDetectionRoutes(statusMcDetectionRoutes, publisherStatusMcDetectionHandler)

}

func configureStatusMcDetectionRoutes(api *gin.RouterGroup, handler *handler.PublisherStatusMcDetectionHandler) {
	api.POST("/create", handler.CreateQueueStatusMcDetection)
}
