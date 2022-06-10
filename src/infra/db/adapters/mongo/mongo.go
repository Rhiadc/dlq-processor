package mongo

import (
	"context"
	"log"

	"github.com/dlqProcessor/src/infra"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongoTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

type MongoAdapter struct {
	config *infra.Config
	Client *mongo.Client
}

func NewMongoClient(ctx context.Context, config *infra.Config) *MongoAdapter {
	opts := options.Client()
	opts.SetMonitor(mongoTrace.NewMonitor())
	opts.ApplyURI(config.MongoURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &MongoAdapter{Client: client, config: config}
}

func (ref MongoAdapter) GetDatabase() *mongo.Database {
	return ref.Client.Database(ref.config.MongoDBName)
}

func (ref MongoAdapter) Close() error {
	return ref.Client.Disconnect(context.Background())
}
