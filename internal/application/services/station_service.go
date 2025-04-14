package services

import (
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/interfaces"
)

type TrainService struct {
	repo interfaces.StationTypeRepository
}

func NewTrainService(repo interfaces.StationTypeRepository) *TrainService {
	return &TrainService{repo: repo}
}

func (s *TrainService) CreateTrainStation(station entities.TrainStation) (entities.TrainStation, error) {
	return s.repo.CreateTrainStation(station)
}

func (s *TrainService) BulkCreateTrainStation(stations []entities.TrainStation) ([]entities.TrainStation, error) {
	return s.repo.BulkCreateTrainStation(stations)
}

func (s *TrainService) UpdateTrainStation(station entities.TrainStation) (entities.TrainStation, error) {
	return s.repo.UpdateTrainStation(station)
}

func (s *TrainService) DeleteTrainStation(code string) error {
	return s.repo.DeleteTrainStation(code)
}

func (s *TrainService) GetTrainStations(filters map[string]interface{}) ([]entities.TrainStation, error) {
	return s.repo.GetTrainStations(filters)
}

func (s *TrainService) GetTrainStationByCode(code string) (entities.TrainStation, error) {
	return s.repo.GetTrainStationByCode(code)
}

// TrainStationType methods
func (s *TrainService) CreateTrainStationType(stationType entities.StationType) (entities.StationType, error) {
	return s.repo.CreateTrainStationType(stationType)
}

func (s *TrainService) UpdateTrainStationType(stationType entities.StationType) (entities.StationType, error) {
	return s.repo.UpdateTrainStationType(stationType)
}

func (s *TrainService) DeleteTrainStationType(code string) error {
	return s.repo.DeleteTrainStationType(code)
}

func (s *TrainService) GetTrainStationTypes() ([]entities.StationType, error) {
	return s.repo.GetTrainStationTypes()
}
