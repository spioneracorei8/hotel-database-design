package server

import (
	"fmt"
	"hotel-database-design/logger"
	"hotel-database-design/route"
	_hotel_handler "hotel-database-design/services/hotel/handler"
	_hotel_repo "hotel-database-design/services/hotel/repostitory"
	_hotel_us "hotel-database-design/services/hotel/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	PSQL_CONNECTION string
	PORT            string
}

func connectDatabase(PSQL_CONNECTION string) *gorm.DB {
	gormLogger := logger.GormLogger()
	database, err := gorm.Open(postgres.Open(PSQL_CONNECTION), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
		return nil
	}
	return database
}

func (s *Server) Start() {
	app := gin.New()
	database := connectDatabase(s.PSQL_CONNECTION)

	//==============================================================
	// # Repositoryies
	//==============================================================
	hotelRepo := _hotel_repo.NewHotelRepoPsqlImpl(database)

	//==============================================================
	// # Usecases
	//==============================================================
	hotelUs := _hotel_us.NewHotelUsecaseImpl(hotelRepo)

	//==============================================================
	// # Handlers
	//==============================================================

	hotelHandler := _hotel_handler.NewHotelHandlerImpl(hotelUs)

	//==============================================================
	// # Routes
	//==============================================================
	app.GET("/", func(g *gin.Context) {
		g.JSON(http.StatusOK, "Hello World!")
	})
	api := route.NewRoute(app)
	api.HotelRoutes(hotelHandler)

	if err := app.Run(fmt.Sprintf(":%s", s.PORT)); err != nil {
		logrus.Fatalf("Error starting server: %v", err)
	}
}
