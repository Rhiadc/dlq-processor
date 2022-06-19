package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dlqProcessor/src/infra"
	"github.com/dlqProcessor/src/infra/db/adapters/mongo"
	"github.com/dlqProcessor/src/infra/db/adapters/mongo/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	config := infra.NewConfig()
	mongoAdapter := mongo.NewMongoClient(context.Background(), config)
	DLQRepo := repository.NewDLQRepository(mongoAdapter.GetDatabase())

	for i := 0; i < 20; i++ {
		message := repository.DLQRecord{Date: time.Now().AddDate(0, 0, -3), Msg: `{"json": "teste"}`, Processed: false}
		_ = DLQRepo.InsertMessage(message)
	}

	g := gin.Default()

	g.GET("/messages/:a", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		limit, _ := strconv.ParseUint(ctx.DefaultQuery("limit", "10"), 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}
		if response, err := DLQRepo.GetAllMessages(a, limit); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": (response)})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
