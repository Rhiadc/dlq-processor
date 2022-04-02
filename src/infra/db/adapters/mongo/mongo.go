package mongo

import (
	"context"
	"github.com/dlqProcessor/src/infra"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongoTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

type MongoAdapter struct {
	config *infra.Config
	client *mongo.Client
}

func NewMongoClient(ctx context.Context, config *infra.Config) *MongoAdapter {
	opts := options.Client()
	opts.SetMonitor(mongoTrace.NewMonitor())
	opts.ApplyURI(config.MongoURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &MongoAdapter{client: client, config: config}
}

func (ref MongoAdapter) GetDatabase() *mongo.Database {
	return ref.client.Database(ref.config.MongoDBName)
}

func (ref MongoAdapter) Close() error {
	return ref.client.Disconnect(context.Background())
}
