package services

import (
	"fmt"

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

func (s *TrainService) DeleteTrainStation(id uint) error {
	return s.repo.DeleteTrainStation(id)
}

func (s *TrainService) GetTrainStations(filters map[string]interface{}) ([]entities.TrainStation, error) {
	return s.repo.GetTrainStations(filters)
}

func (s *TrainService) GetTrainStationByCode(id uint) (entities.TrainStation, error) {
	return s.repo.GetTrainStationById(id)
}

// TrainStationType methods
func (s *TrainService) CreateTrainStationType(stationType entities.StationType) (entities.StationType, error) {
	// Validate required fields
	if stationType.Code == "" {
		return entities.StationType{}, fmt.Errorf("station type code is required")
	}
	if stationType.Name == "" {
		return entities.StationType{}, fmt.Errorf("station type name is required")
	}
	if stationType.ModifyBy == 0 {
		return entities.StationType{}, fmt.Errorf("modifyBy (user ID) is required")
	}

	// Call the repository to create the station type
	createdStationType, err := s.repo.CreateTrainStationType(stationType)
	if err != nil {
		return entities.StationType{}, fmt.Errorf("failed to create station type: %w", err)
	}

	return createdStationType, nil
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
