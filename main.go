package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Doomann/gohome/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	apiPort = 8888
	number  = iota
)

func main() {
	client, ctx, cancel, err := connectToMongo("mongodb://localhost:27018")
	if err != nil {
		panic(err)
	}
	defer closeMongo(client, ctx, cancel)
	checkMongoConnection()
	startWebServer(apiPort)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	controllers.RegisterControllers()

	fmt.Println("Server is up n runnin' on port", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	return port, nil
}

func connectToMongo(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, err
}

func closeMongo(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func checkMongoConnection(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connected to mongo successfully")
	return nil
}
