package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Hotel struct {
	ID               *uuid.UUID `gorm:"primaryKey;index;type:uuid;default:uuid_generate_v4();not null" json:"id"`
	HotelName        string     `gorm:"size:255;not null" json:"hotel_name"`
	HotelDescription *string    `gorm:"size:255" json:"hotel_description"`

	Address           string  `gorm:"type:text;not null" json:"address"`
	ProvinceNameTH    string  `gorm:"size:255;not null" json:"province_name_th"`
	ProvinceNameEN    string  `gorm:"size:255;not null" json:"province_name_en"`
	DistrictNameTH    string  `gorm:"size:255;not null" json:"district_name_th"`
	DistrictNameEN    string  `gorm:"size:255;not null" json:"district_name_en"`
	SubdistrictNameTH string  `gorm:"size:255;not null" json:"subdistrict_name_th"`
	SubdistrictNameEN string  `gorm:"size:255;not null" json:"subdistrict_name_en"`
	TelephoneNumber   string  `gorm:"size:10;not null" json:"telephone_number"`
	MobilephoneNumber *string `gorm:"size:10" json:"mobilephone_number"`
	Email             string  `gorm:"size:255;not null" json:"email"`
	LineID            *string `gorm:"size:255" json:"line_id"`
	FacebookID        *string `gorm:"size:255" json:"facebook_id"`
	InstagramID       *string `gorm:"size:255" json:"instagram_id"`

	Status string `gorm:"type:hotel_status;default:'OPEN';not null" json:"status"`

	CreatedBy string     `gorm:"size:50;not null" json:"created_by"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedBy *string    `gorm:"size:50" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedBy *string    `gorm:"size:50" json:"deleted_by"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Hotel) TableName() string {
	return "hotel"
}

func (h *Hotel) GenUUID() {
	id, _ := uuid.NewV4()
	h.ID = &id
}
