package services

import "gorm.io/gorm"

type (
	IncidentService struct {
		db *gorm.DB
	}

	Incident struct {
		Object     string       `json:"object"`
		ObjectData []ObjectData `json:"object_data"`
		Message    string       `json:"message"`
		Date       string       `json:"date"`
	}

	ObjectData struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	IncidentWrapper struct {
		Service  string   `json:"service"`
		Incident Incident `json:"incident"`
	}

	IIncidentService interface {
		ProccessNewIncident(data IncidentWrapper) (bool, error)
	}
)

func NewIncidentService(db *gorm.DB) *IncidentService {
	return &IncidentService{db: db}
}

func (inc *IncidentService) ProccessNewIncident(data IncidentWrapper) (bool, error) {
	return true, nil
}
