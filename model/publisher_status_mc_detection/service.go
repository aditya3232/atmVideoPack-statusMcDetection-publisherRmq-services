package publisher_status_mc_detection

type Service interface {
	CreateQueueStatusMcDetection(input RmqPublisherStatusMcDetectionInput) (RmqPublisherStatusMcDetection, error)
}

type service struct {
	statusMcDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// public message to rmq
func (s *service) CreateQueueStatusMcDetection(input RmqPublisherStatusMcDetectionInput) (RmqPublisherStatusMcDetection, error) {

	newRmqPublisherStatusMcDetection := RmqPublisherStatusMcDetection{
		Tid:           input.Tid,
		DateTime:      input.DateTime,
		StatusSignal:  input.StatusSignal,
		StatusStorage: input.StatusStorage,
		StatusRam:     input.StatusRam,
		StatusCpu:     input.StatusCpu,
	}

	newStatusMcDetection, err := s.statusMcDetectionRepository.CreateQueueStatusMcDetection(newRmqPublisherStatusMcDetection)
	if err != nil {
		return newStatusMcDetection, err
	}

	return newStatusMcDetection, nil
}
