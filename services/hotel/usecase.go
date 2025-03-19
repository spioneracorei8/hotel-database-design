package hotel

import (
	"hotel-database-design/models"
	"sync"
)

type HotelUsecase interface {
	FetchHotels(paginator *models.Pagination, args *sync.Map) ([]*models.Hotel, *models.Pagination, error)
	CreateHotel(hotel *models.Hotel) error
}
