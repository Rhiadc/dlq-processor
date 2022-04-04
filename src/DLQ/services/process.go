package services

import (
	"github.com/dlqProcessor/src/infra/db/adapters/mongo/repository"
)

type ProcessService struct {
	DLQRepo *repository.DLQRepository
	//kafka writer (TOPIC)
}

func NewProcessService(DLQRepo *repository.DLQRepository) *ProcessService {
	return &ProcessService{DLQRepo: DLQRepo}
}

func (ref ProcessService) Process(dateFrom int) error {
	dlqMessages, err := ref.DLQRepo.GetMessagesByDateRange(dateFrom)
	if err != nil {
		return err
	}
	for _, msg := range dlqMessages {
		result, err := ref.DLQRepo.GetMessageByDate(msg.Date)
		if err != nil {
			return err
		}
		if result.Processed == false {
			//write messate to topic
		}

		err = ref.DLQRepo.SetProcessedMessage(msg.Date)
		if err != nil {
			return err
		}
	}
	return nil
}
