package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Mensurui/movieReservation/internal/domain"
	"github.com/Mensurui/movieReservation/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	userService service.UserService
}

func NewUserHandlers(userService service.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (uh *UserHandlers) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := uh.userService.Register(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (uh *UserHandlers) GetMovie(c *gin.Context) {
	dateStr := c.Query("date")
	timeStr := c.Query("time")
	if dateStr == "" || timeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date and time query parameters are required"})
		return
	}

	dateTimeStr := fmt.Sprintf("%s %s", dateStr, timeStr)
	layout := "2006-01-02 15:04"

	parsedShowtime, err := time.ParseInLocation(layout, dateTimeStr, time.Local)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid date or time format. Expected date as YYYY-MM-DD and time as HH:MM. Received: date='%s', time='%s'", dateStr, timeStr),
		})
		return
	}

	movie, err := uh.userService.GetMovie(c.Request.Context(), parsedShowtime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}

func (uh *UserHandlers) ReserveSeat(c *gin.Context) {
	userID := c.Param("id")
	moviepremierID := c.Query("moviepid")
	if userID == "" || moviepremierID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date and time query parameters are required"})
		return
	}
	mid, _ := strconv.Atoi(moviepremierID)
	uid, _ := strconv.Atoi(userID)

	err := uh.userService.ReserveSeat(c.Request.Context(), mid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}
