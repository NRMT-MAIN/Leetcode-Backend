package config

import (
	"context"
	"fmt"
	"leetcode/config/env"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func CreateClient() (*mongo.Client , error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
    defer cancel()

	uri := env.GetString("MONGO_DB_URI"  , "mongodb://localhost:27017")
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))


    if err != nil {
		fmt.Println("Connection to MongoDB failed!" , err)
        return nil, err
    }
	fmt.Println("MongoDB connected successfully")
    return client , nil
}


