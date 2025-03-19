package handler

import (
	"hotel-database-design/constants"
	"hotel-database-design/models"
	"hotel-database-design/services/hotel"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type hotelHandler struct {
	hotelUsecase hotel.HotelUsecase
}

func NewHotelHandlerImpl(hotelUsecase hotel.HotelUsecase) hotel.HotelHandler {
	return &hotelHandler{
		hotelUsecase: hotelUsecase,
	}
}

func (r *hotelHandler) FetchHotels(g *gin.Context) {
	var paginator = new(models.Pagination)
	var args = new(sync.Map)
	page, err := strconv.Atoi(g.Query("page"))
	if err != nil {
		page = 1
	}
	perPage, err := strconv.Atoi(g.Query("per_page"))
	if err != nil {
		perPage = 10
	}
	hotelName := g.Query("hotel_name")
	if hotelName != "" {
		args.Store("hotel_name", hotelName)
	}
	paginator.Page = page
	paginator.PerPage = perPage
	hotels, pagination, err := r.hotelUsecase.FetchHotels(paginator, args)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"status": constants.FAILED_STATUS,
			"error":  err.Error(),
		})
	}
	g.JSON(http.StatusOK, gin.H{
		"hotels":     hotels,
		"pagination": pagination,
		"status":     constants.SUCCESS_STATUS,
	})
}

func (r *hotelHandler) CreateHotel(g *gin.Context) {
	var hotel = new(models.Hotel)
	if err := g.ShouldBind(&hotel); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status": constants.FAILED_STATUS,
			"error":  err.Error(),
		})
	}
	if err := r.hotelUsecase.CreateHotel(hotel); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"status": constants.FAILED_STATUS,
			"error":  err.Error(),
		})
	}
	g.JSON(http.StatusOK, gin.H{
		"status": constants.SUCCESS_STATUS,
	})
}
