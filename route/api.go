package route

import (
	"hotel-database-design/services/hotel"

	"github.com/gin-gonic/gin"
)

type route struct {
	g *gin.Engine
}

func NewRoute(g *gin.Engine) *route {
	return &route{
		g: g,
	}
}

func (r *route) HotelRoutes(handler hotel.HotelHandler) {
	api := r.g.Group("/api")

	api.GET("/v1/hotels", handler.FetchHotels)
	api.POST("/v1/hotel", handler.CreateHotel)
}
