package usecase

import (
	"hotel-database-design/models"
	"hotel-database-design/services/hotel"
	"sync"
)

type hotelUsecase struct {
	hotelRepo hotel.HotelRepository
}

func NewHotelUsecaseImpl(hotelRepo hotel.HotelRepository) hotel.HotelUsecase {
	return &hotelUsecase{
		hotelRepo: hotelRepo,
	}
}

func (r *hotelUsecase) FetchHotels(paginator *models.Pagination, args *sync.Map) ([]*models.Hotel, *models.Pagination, error) {
	return r.hotelRepo.FetchHotels(paginator, args)
}
func (r *hotelUsecase) CreateHotel(hotel *models.Hotel) error {
	hotel.GenUUID()
	return r.hotelRepo.CreateHotel(hotel)
}
