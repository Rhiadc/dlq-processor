package main

import (
	"context"
	"fmt"
	"github.com/dlqProcessor/src/infra"
	"github.com/dlqProcessor/src/infra/db/adapters/mongo"
	"github.com/dlqProcessor/src/infra/db/adapters/mongo/repository"
	"time"
)

func main() {
	//@TODO
	config := infra.NewConfig()
	mongoAdapter := mongo.NewMongoClient(context.Background(), config)
	DLQRepo := repository.NewDLQRepository(mongoAdapter.GetDatabase())
	message := repository.DLQRecord{Date: time.Now().AddDate(0, 0, -3), Msg: `{"json": "teste"}`, Processed: false}
	err := DLQRepo.InsertMessage(message)

	lastMessages, err := DLQRepo.GetMessagesByDateRange(-4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Antes excluir DLQ...", lastMessages[0])
	err = DLQRepo.SetProcessedMessage(lastMessages[0].Date)
	if err != nil {
		fmt.Println(err)
	}
	lastMessages, _ = DLQRepo.GetAllMessages()
	fmt.Println("Processed messages: ", lastMessages)
}
