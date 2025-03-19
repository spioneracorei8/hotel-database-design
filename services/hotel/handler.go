package hotel

import (
	"github.com/gin-gonic/gin"
)

type HotelHandler interface {
	FetchHotels(g *gin.Context)
	CreateHotel(g *gin.Context)
}
