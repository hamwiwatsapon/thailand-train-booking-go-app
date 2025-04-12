package interfaces

import (
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
)

type TrainRepository interface {
	// TrainStation
	CreateTrainStation(station entities.TrainStation) (entities.TrainStation, error)
	BulkCreateTrainStation(stations []entities.TrainStation) ([]entities.TrainStation, error)
	UpdateTrainStation(station entities.TrainStation) (entities.TrainStation, error)
	DeleteTrainStation(code string) error
	GetTrainStations(filters map[string]interface{}) ([]entities.TrainStation, error)
	GetTrainStationByCode(code string) (entities.TrainStation, error)

	// TrainStationType
	CreateTrainStationType(stationType entities.StationType) (entities.StationType, error)
	UpdateTrainStationType(stationType entities.StationType) (entities.StationType, error)
	DeleteTrainStationType(code string) error
	GetTrainStationTypes() ([]entities.StationType, error)
}
