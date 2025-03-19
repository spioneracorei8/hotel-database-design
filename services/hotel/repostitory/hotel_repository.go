package repostitory

import (
	"hotel-database-design/models"
	"hotel-database-design/services/hotel"
	"strings"
	"sync"

	"gorm.io/gorm"
)

type hotelRepository struct {
	db *gorm.DB
}

func NewHotelRepoPsqlImpl(db *gorm.DB) hotel.HotelRepository {
	return &hotelRepository{
		db: db,
	}
}

func (r *hotelRepository) FetchHotels(paginator *models.Pagination, args *sync.Map) ([]*models.Hotel, *models.Pagination, error) {
	var hotels []*models.Hotel
	var where string
	conds, valArgs := r.whereFetchHotels(args)
	if len(conds) > 0 {
		where = "deleted_at IS NULL " + " AND " + strings.Join(conds, " AND ")
	}
	if err := r.db.
		Scopes(models.Paginate(hotels, where, valArgs, paginator, r.db)).
		Where(where, valArgs...).
		Order("updated_at DESC").
		Find(&hotels).
		Error; err != nil {
		return nil, nil, err
	}
	return hotels, paginator, nil
}

func (r *hotelRepository) whereFetchHotels(args *sync.Map) ([]string, []interface{}) {
	var (
		conds   = make([]string, 0)
		valArgs = make([]interface{}, 0)
	)
	if v, ok := args.Load("hotel_name"); ok {
		if v != nil && v != "" {
			sql := `hotel_name LIKE ?`
			conds = append(conds, sql)
			valArgs = append(valArgs, "%"+v.(string)+"%")
		}
	}

	return conds, valArgs
}

func (r *hotelRepository) CreateHotel(hotel *models.Hotel) error {
	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.
		Create(&hotel).
		Error; err != nil {
		return err
	}
	return tx.Commit().Error
}
