package router

import (
	"strings"
	"time"

	"github.com/adhyttungga/go-crud-gin-swagger/config"
	tags_delivery "github.com/adhyttungga/go-crud-gin-swagger/delivery/tags"
	tags_repository "github.com/adhyttungga/go-crud-gin-swagger/repository/tags"
	tags_usecase "github.com/adhyttungga/go-crud-gin-swagger/usecase/tags"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(DB *gorm.DB) *gin.Engine {
	router := gin.Default()
	allowOrigins := strings.Split(config.Config.Origin.AllowOrigin, ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}))

	// Validator
	validate := validator.New()

	// Repository
	tagsRepository := tags_repository.NewTagsRepository(DB)

	// Usecase
	tagsUsecase := tags_usecase.NewTagsUsecaseImpl(tagsRepository, validate)

	// Delivery
	tagsDelivery := tags_delivery.NewTagsDeliveryImpl(tagsUsecase)

	// Add Swagger Router
	if config.Config.GinMode == "debug" {
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	}

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	{
		tagsRouter.POST("", tagsDelivery.Create)
		tagsRouter.GET("", tagsDelivery.FindAll)
		tagsRouter.GET("/:tagId", tagsDelivery.FindById)
	}

	return router
}
