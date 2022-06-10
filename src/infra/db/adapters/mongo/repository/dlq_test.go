package repository_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/dlqProcessor/src/infra"
	cmongo "github.com/dlqProcessor/src/infra/db/adapters/mongo"
	"github.com/dlqProcessor/src/infra/db/adapters/mongo/repository"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

const MONGO_INITDB_ROOT_USERNAME = "root"
const MONGO_INITDB_ROOT_PASSWORD = "password"

func TestMain(m *testing.M) {
	// Setup
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	environmentVariables := []string{
		"MONGO_INITDB_ROOT_USERNAME=" + MONGO_INITDB_ROOT_USERNAME,
		"MONGO_INITDB_ROOT_PASSWORD=" + MONGO_INITDB_ROOT_PASSWORD,
	}
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "5.0",
		Env:        environmentVariables,
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		var err error
		db, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://%s:%s@localhost:%s", MONGO_INITDB_ROOT_USERNAME, MONGO_INITDB_ROOT_PASSWORD, resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return db.Ping(context.TODO(), nil)
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Run tests
	exitCode := m.Run()

	// Teardown
	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// disconnect mongodb client
	if err = db.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	// Exit
	os.Exit(exitCode)
}

func TestAddTodo(t *testing.T) {
	DLQRecord := repository.DLQRecord{
		Date:      time.Now(),
		Msg:       "something",
		Processed: true,
	}

	client := cmongo.NewMongoClient(context.TODO(), infra.NewConfig())
	client.Client = db

	todos := repository.NewDLQRepository(client.GetDatabase())
	// add tod
	err := todos.InsertMessage(DLQRecord)
	// assert error is nil
	assert.Nil(t, err)
}
