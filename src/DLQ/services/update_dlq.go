package services

import (
	"github.com/dlqProcessor/src/infra/db/adapters/mongo/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateDLQService struct {
	DLQRepo *repository.DLQRepository
}

func NewUpdateDLQService(DLQRepo *repository.DLQRepository) *UpdateDLQService {
	return &UpdateDLQService{DLQRepo: DLQRepo}
}

func (ref UpdateDLQService) UpdateDLQ(dlqMessage repository.DLQRecord) error {
	message, err := ref.DLQRepo.GetMessageByDate(dlqMessage.Date)
	if err == mongo.ErrNoDocuments {
		err := ref.DLQRepo.InsertMessage(dlqMessage)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	if message.Processed == false {
		err = ref.DLQRepo.SetProcessedMessage(message.Date)
		if err != nil {
			return err
		}
	}

	return nil
}
