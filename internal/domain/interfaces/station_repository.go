package interfaces

import (
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
)

type StationTypeRepository interface {
	// TrainStation
	CreateTrainStation(station entities.TrainStation) (entities.TrainStation, error)
	BulkCreateTrainStation(stations []entities.TrainStation) ([]entities.TrainStation, error)
	UpdateTrainStation(station entities.TrainStation) (entities.TrainStation, error)
	DeleteTrainStation(id uint) error
	GetTrainStations(filters map[string]interface{}) ([]entities.TrainStation, error)
	GetTrainStationById(id uint) (entities.TrainStation, error)

	// TrainStationType
	CreateTrainStationType(stationType entities.StationType) (entities.StationType, error)
	UpdateTrainStationType(stationType entities.StationType) (entities.StationType, error)
	DeleteTrainStationType(code string) error
	GetTrainStationTypes() ([]entities.StationType, error)
}
