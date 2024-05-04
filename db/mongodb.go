package db

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/phonghaido/golang-mongodb/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongoDB(ctx context.Context) mongo.Client {
	file, err := os.Open("config.txt")
	if err != nil {
		log.Println("Error opening config file (config.txt)")
	}
	defer file.Close()

	config := make(map[string]string)

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		lines := strings.Split(string(buffer[:n]), "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(strings.Trim(parts[1], `"`))
				config[key] = value
			}
		}
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config["MONGODB_URI"]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return *client
}

func FetchRecords(ctx context.Context, client *mongo.Client, pipeline bson.A) []data.MongoDBRecord {
	coll := client.Database("getircase-study").Collection("records")
	defer client.Disconnect(ctx)
	cur, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var records []data.MongoDBRecord
	if err := cur.All(ctx, &records); err != nil {
		log.Fatal(err)
	}

	return records
}
