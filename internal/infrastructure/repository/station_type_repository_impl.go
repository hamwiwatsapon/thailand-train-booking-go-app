package repository

import (
	"fmt"

	"github.com/hamwiwatsapon/train-booking-go/internal/application/utils"
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/entities"
	"github.com/hamwiwatsapon/train-booking-go/internal/domain/interfaces"
	"gorm.io/gorm"
)

func NewTrainRepository(db *gorm.DB) interfaces.StationTypeRepository {
	return &trainRepositoryImpl{db: db}
}

type trainRepositoryImpl struct {
	db *gorm.DB
}

// CreateTrainStation implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) CreateTrainStation(station entities.TrainStation) (entities.TrainStation, error) {
	tx := t.db.Begin()
	if err := tx.Error; err != nil {
		return entities.TrainStation{}, err
	}

	if err := tx.Where("code = ?", station.Code).First(&entities.User{}).Error; err == nil {
		tx.Rollback()
		return entities.TrainStation{}, fmt.Errorf("train station code %s already exists", station.Code)
	}

	if err := tx.Create(&station).Error; err != nil {
		tx.Rollback()
		return entities.TrainStation{}, err
	}

	return station, tx.Commit().Error
}

// DeleteTrainStation implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) DeleteTrainStation(code string) error {
	// Start a transaction
	tx := t.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	// Check if the user exists and delete in one step
	if err := tx.Where("code = ?", code).Delete(&entities.TrainStation{}).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("train station code %s not found", code)
		}
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

// GetTrainStations implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) GetTrainStations(filters map[string]interface{}) ([]entities.TrainStation, error) {
	var trainStations []entities.TrainStation

	// Apply filters to the query
	query := t.db.Model(&entities.TrainStation{})
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	if err := query.Find(&trainStations).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch train stations: %w", err)
	}

	return trainStations, nil
}

// UpdateTrainStation implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) UpdateTrainStation(station entities.TrainStation) (entities.TrainStation, error) {
	// Check if the user exists
	var existingTrainStation entities.TrainStation
	if err := t.db.Where("code = ?", station.Code).First(&existingTrainStation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.TrainStation{}, fmt.Errorf("train station with code %s not found", station.Code)
		}
		return entities.TrainStation{}, fmt.Errorf("failed to fetch train station with code %s: %w", station.Code, err)
	}

	// Update the user
	if err := t.db.Model(&existingTrainStation).Updates(station).Error; err != nil {
		return entities.TrainStation{}, fmt.Errorf("failed to update train station with code %s: %w", station.Code, err)
	}

	return existingTrainStation, nil
}

// GetTrainStationByCode implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) GetTrainStationByCode(code string) (entities.TrainStation, error) {
	var trainStation entities.TrainStation

	if err := t.db.Where("code = ?", code).First(&entities.User{}).Error; err != nil {
		return entities.TrainStation{}, fmt.Errorf("failed to fetch train station with code %s: %w", code, err)
	}

	return trainStation, nil
}

// BulkCreateTrainStation implements interfaces.TrainRepository.
// This method creates multiple train stations in bulk. It first checks if any of the provided train station codes already exist in the database.
// If any of the codes already exist, it rolls back the transaction and returns an error.
func (t *trainRepositoryImpl) BulkCreateTrainStation(stations []entities.TrainStation) ([]entities.TrainStation, error) {
	tx := t.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	existingCodes := make(map[string]struct{})
	for _, station := range stations {
		existingCodes[station.Code] = struct{}{}
	}

	var existingStations []entities.TrainStation
	if err := tx.Where("code IN ?", utils.GetKeys(existingCodes)).Find(&existingStations).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, existingStation := range existingStations {
		if _, exists := existingCodes[existingStation.Code]; exists {
			tx.Rollback()
			return nil, fmt.Errorf("train station code %s already exists", existingStation.Code)
		}
	}

	if err := tx.Create(&stations).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return stations, tx.Commit().Error
}

// CreateTrainStationType implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) CreateTrainStationType(stationType entities.StationType) (entities.StationType, error) {
	tx := t.db.Begin()
	if err := tx.Error; err != nil {
		return entities.StationType{}, err
	}

	if err := tx.Where("code = ?", stationType.Code).First(&entities.TrainStation{}).Error; err == nil {
		tx.Rollback()
		return entities.StationType{}, fmt.Errorf("train station type code %s already exists", stationType.Code)
	}

	if err := tx.Create(&stationType).Error; err != nil {
		tx.Rollback()
		return entities.StationType{}, err
	}

	return stationType, tx.Commit().Error
}

// UpdateTrainStationType implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) UpdateTrainStationType(stationType entities.StationType) (entities.StationType, error) {
	// Check if the user exists
	var existingTrainStationType entities.StationType
	if err := t.db.Where("code = ?", stationType.Code).First(&existingTrainStationType).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.StationType{}, fmt.Errorf("train station type with code %s not found", stationType.Code)
		}
		return entities.StationType{}, fmt.Errorf("failed to fetch train station type with code %s: %w", stationType.Code, err)
	}

	// Update the train station type
	if err := t.db.Model(&existingTrainStationType).Updates(stationType).Error; err != nil {
		return entities.StationType{}, fmt.Errorf("failed to update train station type with code %s: %w", stationType.Code, err)
	}

	return existingTrainStationType, nil
}

// DeleteTrainStationType implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) DeleteTrainStationType(code string) error {
	// Start a transaction
	tx := t.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	// Check if the user exists and delete in one step
	if err := tx.Where("code = ?", code).Delete(&entities.StationType{}).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("train station type code %s not found", code)
		}
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

// GetTrainStationTypes implements interfaces.TrainRepository.
func (t *trainRepositoryImpl) GetTrainStationTypes() ([]entities.StationType, error) {
	var trainStationTypes []entities.StationType

	if err := t.db.Find(&trainStationTypes).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch train station types: %w", err)
	}

	return trainStationTypes, nil
}
