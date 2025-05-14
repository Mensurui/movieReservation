package http

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Mensurui/movieReservation/internal/data/postgres"
	"github.com/Mensurui/movieReservation/internal/service"
	"github.com/Mensurui/movieReservation/internal/transport/http/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//TODO: Put the authentication here router.Use()
	adminRepo := postgres.NewPostgresAdminRepository(db)
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handlers.NewAdminHandlers(*adminService)

	userRepo := postgres.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandlers(*userService)

	apiV1 := router.Group("/v1")

	{
		adminGroup := apiV1.Group("/admin")
		{
			adminGroup.POST("/movie/add", adminHandler.AddMovie)
			adminGroup.PATCH("/movie/update", adminHandler.UpdateMovie)
			adminGroup.DELETE("/movie/delete/:id", adminHandler.DeleteMovie)
			adminGroup.GET("/movies", adminHandler.GetMovie)

			adminGroup.POST("/theater/add", adminHandler.AddTheater)
			adminGroup.GET("/theater/get", adminHandler.GetTheaterCapacity)

			adminGroup.POST("/movie-premier/add", adminHandler.AddMoviePremier)

		}
	}

	{
		userGroup := apiV1.Group("/user")
		{
			userGroup.POST("/register", userHandler.Register)
			userGroup.GET("/movie", userHandler.GetMovie)
			userGroup.POST("/reserve-seat/:id", userHandler.ReserveSeat)
		}
	}

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	return router

}
