package publisher_status_mc_detection

type Service interface {
	CreateQueueStatusMcDetection(input PublisherStatusMcInput) (StatusMcDetection, error)
}

type service struct {
	statusMcDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// public message to rmq
func (s *service) CreateQueueStatusMcDetection(input PublisherStatusMcInput) (StatusMcDetection, error) {
	_, err := s.statusMcDetectionRepository.CreateQueueStatusMcDetection(input)
	if err != nil {
		return StatusMcDetection{}, err
	}

	// Convert input to StatusMcDetection
	statusMcDetection := StatusMcDetection(input)

	// return the converted input as StatusMcDetection
	return statusMcDetection, nil
}
