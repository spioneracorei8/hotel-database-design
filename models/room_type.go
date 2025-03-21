package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type RoomType struct {
	ID       *uuid.UUID `gorm:"primaryKey;index;type:uuid;default:uuid_generate_v4()" json:"id"`
	RoomType string     `gorm:"size:100;not null" json:"room_type"`

	IsActive bool `gorm:"default:true;not null" json:"is_active"`

	CreatedBy string     `gorm:"size:50;not null" json:"created_by"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedBy *string    `gorm:"size:50" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedBy *string    `gorm:"size:50" json:"deleted_by"`
	DeletedAt *time.Time `json:"deleted_at"`

	HotelRoomType *HotelRoomType `gorm:"foreignKey:RoomTypeID" json:"hotel_room_type"`
}

func (RoomType) TableName() string {
	return "room_type"
}
func (rt *RoomType) GenUUID() {
	id, _ := uuid.NewV4()
	rt.ID = &id
}
