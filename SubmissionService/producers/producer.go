package producers

import (
	"Submission_Service/config/redis"
	"context"
	"encoding/json"
	"fmt"
)

func ProduceJob(queueName string, payload interface{}) error {
	client := redis.CreateRedisClient()
	if client == nil {
		return fmt.Errorf("failed to create Redis client")
	}
	defer client.Close()

	ctx := context.Background()
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	err = client.LPush(ctx, queueName, data).Err()
	if err != nil {
		return fmt.Errorf("failed to push job to Redis: %w", err)
	}

	return nil
}