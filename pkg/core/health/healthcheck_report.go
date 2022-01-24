package health

import "go-healthcheck/pkg/datamodel"

type service struct {
}

// NewService returns address service which implements Service interface
func NewService() Service {
	return &service{}
}

// Service defines address usecases (see implementation at usecase_*.go)
type Service interface {
	GetHealthCheckReport() (resp *datamodel.HealthResponse, err error)
}
