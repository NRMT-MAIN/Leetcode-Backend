package config

import (
	"Submission_Service/config/env"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func CreateClient() (*mongo.Client, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := env.GetString("MONGO_DB_URI", "mongodb://localhost:27017")
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        fmt.Println("Connection to MongoDB failed:", err)
        return nil, err
    }

    // ✅ Verify connection
    if err := client.Ping(ctx, nil); err != nil {
        fmt.Println("MongoDB ping failed:", err)
        return nil, err
    }

    fmt.Println("MongoDB connected successfully")
    return client, nil
}
