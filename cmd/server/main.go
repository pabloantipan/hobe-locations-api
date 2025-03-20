package main

import (
	"context"
	"log"
	"net/http"

	"github.com/pabloantipan/hobe-locations-api/config"
	"github.com/pabloantipan/hobe-locations-api/internal/bussines"
	"github.com/pabloantipan/hobe-locations-api/internal/cloud"
	"github.com/pabloantipan/hobe-locations-api/internal/middleware"
	"github.com/pabloantipan/hobe-locations-api/internal/repositories/datastore"
	"github.com/pabloantipan/hobe-locations-api/internal/repositories/storage"
	"github.com/pabloantipan/hobe-locations-api/internal/services"

	"github.com/pabloantipan/hobe-locations-api/internal/handlers"

	"github.com/gin-gonic/gin"

	_ "github.com/pabloantipan/hobe-locations-api/docs/swagger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Player API
// @version         1.0
// @description     API Server for Player Management
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {
	ctx := context.Background()
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	logger, err := cloud.NewCloudLogger(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Cloud Logger: %v", err)
	}

	// Initialize Datastore client
	datastoreClient := datastore.NewDatastoreClient(cfg)
	storageClient := storage.NewStorageClient(cfg)

	// Initialize repositories
	pictureRepo := storage.NewPictureRepository(storageClient)
	locationRepo := datastore.NewDatastoreLocationRepository(&ctx, datastoreClient)

	// Initialize services
	pictureService := services.NewPictureService(&pictureRepo)
	locationService := services.NewLocationService(locationRepo)

	// Initialize businesses
	locationBusiness := bussines.NewLocationsBusiness(pictureService, locationService)

	// Initialize middlewares
	requestLoggerMiddleware := middleware.NewRequestLoggerMiddleware(logger)
	responseLoggerMiddleware := middleware.NewResponseLoggerMiddleware(logger)

	// Initialize handlers
	locationHandler := handlers.NewLocationsHandler(&locationBusiness)
	configsHandler := handlers.NewConfigsHandler()

	// Add(rateLimiter.Handle)
	healthHandler := handlers.NewHealthHandler(cfg)

	// Setup router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)

	router.GET("/health", func(c *gin.Context) {
		healthHandler.ServeHTTP(c.Writer, c.Request)
	})

	api := router.Group("locations/api/v1")
	{
		locations := api.Group("")
		{
			locations.Use(requestLoggerMiddleware.HandleFunc())
			locations.Use(responseLoggerMiddleware.HandleFunc())
			locations.POST("", locationHandler.Add)
			locations.GET("", locationHandler.GetThemByEmail)
		}

		markers := api.Group("markers")
		{
			markers.Use(requestLoggerMiddleware.HandleFunc())
			markers.Use(responseLoggerMiddleware.HandleFunc())
			markers.POST("", locationHandler.GetThemByMapSquare)
		}

		configs := api.Group("configs")
		{
			// configs.Use(requestLoggerMiddleware.HandleFunc())
			// configs.Use(responseLoggerMiddleware.HandleFunc())
			configs.GET("point-types", configsHandler.GetPointTypes)
			configs.GET("order-key-options", configsHandler.GetLocationOrderKeys)
		}
	}

	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Starting API Gateway on port %s: %v", cfg.Port, err)
	}
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
