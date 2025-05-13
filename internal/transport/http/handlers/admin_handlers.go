package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Mensurui/movieReservation/internal/domain"
	"github.com/Mensurui/movieReservation/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminHandlers struct {
	adminService service.AdminService
}

func NewAdminHandlers(adminService service.AdminService) *AdminHandlers {
	return &AdminHandlers{
		adminService: adminService,
	}
}

func (ah *AdminHandlers) AddMovie(c *gin.Context) {
	var movie domain.Movie
	if err := c.ShouldBindBodyWithJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	err := ah.adminService.AddMovie(c.Request.Context(), movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	log.Printf("Movies: %v", movie)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ah *AdminHandlers) UpdateMovie(c *gin.Context) {
	var movie domain.Movie
	if err := c.ShouldBindBodyWithJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err := ah.adminService.UpdateMovie(c.Request.Context(), movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (ah *AdminHandlers) DeleteMovie(c *gin.Context) {
	value := c.Param("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	err = ah.adminService.DeleteMovie(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ah *AdminHandlers) GetMovie(c *gin.Context) {
	movies, err := ah.adminService.GetMovie(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func (ah *AdminHandlers) AddTheater(c *gin.Context) {
	var theater domain.Theater
	if err := c.ShouldBindBodyWithJSON(&theater); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	err := ah.adminService.AddTheater(c.Request.Context(), theater)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ah *AdminHandlers) GetTheaterCapacity(c *gin.Context) {
	name := c.Query("name")
	capacity, err := ah.adminService.GetTheaterCapacity(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"capcity": capacity})
}

func (ah *AdminHandlers) AddMoviePremier(c *gin.Context) {
	var moviePremier domain.MoviePremier
	if err := c.ShouldBindBodyWithJSON(&moviePremier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	if err := ah.adminService.AddMoviePremier(c.Request.Context(), moviePremier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
