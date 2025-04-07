package entities

import (
	"time"

	"gorm.io/gorm"
)

// Train represents a train entity.
type Train struct {
	Code            string `json:"code" gorm:"primaryKey;not null;index"`
	Name            string `json:"name" gorm:"not null"`
	FromStationCode string `json:"from" gorm:"not null"` // FK to TrainStation
	ToStationCode   string `json:"to" gorm:"not null"`   // FK to TrainStation
	Time            string `json:"time" gorm:"not null"`
	Price           int    `json:"price" gorm:"not null"`
	Seats           int    `json:"seats" gorm:"not null"`
	AvailableSeats  int    `json:"available_seats" gorm:"not null"`

	TrainTypeCode string    `json:"train_type_code" gorm:"not null"` // FK to TrainType
	TrainType     TrainType `json:"train_type" gorm:"foreignKey:TrainTypeCode;references:Code"`

	FromStation TrainStation `gorm:"foreignKey:FromStationCode;references:Code"`
	ToStation   TrainStation `gorm:"foreignKey:ToStationCode;references:Code"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ModifyBy  uint           `json:"modify_by" gorm:"not null"` // FK to User
	User      User           `gorm:"foreignKey:ModifyBy;references:ID"`
}

// TrainType represents different types of trains.
type TrainType struct {
	Code string `json:"code" gorm:"primaryKey;not null;index"`
	Name string `json:"name" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ModifyBy  uint           `json:"modify_by" gorm:"not null"`
	User      User           `gorm:"foreignKey:ModifyBy;references:ID"`
}

// TrainStation represents a train station.
type TrainStation struct {
	Code string `json:"code" gorm:"primaryKey;not null;uniqueIndex"`
	Name string `json:"name" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ModifyBy  uint           `json:"modify_by" gorm:"not null"`
	User      User           `gorm:"foreignKey:ModifyBy;references:ID"`
}

// StationOrder represents an ordered set of stations for a train route.
type StationOrder struct {
	gorm.Model
	TrainCode string `json:"train_code" gorm:"not null;index"` // FK to Train
	Train     Train  `gorm:"foreignKey:TrainCode;references:Code"`

	Stations []StationOrderDetail `json:"stations" gorm:"foreignKey:StationOrderID"`

	ModifyBy uint `json:"modify_by" gorm:"not null"`
	User     User `gorm:"foreignKey:ModifyBy;references:ID"`
}

// StationOrderDetail represents the details of a station in a train route.
type StationOrderDetail struct {
	gorm.Model
	StationOrderID uint         `json:"station_order_id" gorm:"not null"` // FK to StationOrder
	StationCode    string       `json:"station_code" gorm:"not null"`     // FK to TrainStation
	TrainStation   TrainStation `gorm:"foreignKey:StationCode;references:Code"`

	Order int `json:"order" gorm:"not null"` // Order in the train route

	ModifyBy uint `json:"modify_by" gorm:"not null"`
	User     User `gorm:"foreignKey:ModifyBy;references:ID"`
}
