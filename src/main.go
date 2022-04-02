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
	message := repository.DLQRecord{Date: time.Now().AddDate(0, 0, -3), Msg: `{"json": "teste"}`, Processed: true}
	err := DLQRepo.InsertMessage(message)

	lastMessages, err := DLQRepo.GetMessagesByDateRange(-4)
	if err != nil {
		fmt.Println(err)
	}

	err = DLQRepo.DeleteMessageByDate(lastMessages[0].Date)
	if err != nil {
		fmt.Println(err)
	}

	processed, err := DLQRepo.GetProcessedMessages()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Processing DLQ...")
	fmt.Println("Processed messages: ", processed)
	fmt.Println("last messages: ", processed)
}
