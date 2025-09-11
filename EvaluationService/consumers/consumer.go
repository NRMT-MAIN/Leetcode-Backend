package consumers

import (
	config "Evaluation_Service/config/redis"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

var ctx = context.Background()

type Submission struct {
    ID       string `json:"id"`
    Code     string `json:"code"`
    Language string `json:"language"`
}

func startWorker() {
	client := config.CreateRedisClient()
	if client == nil {
		fmt.Println("Failed to create Redis client")
		return
	}

    for {
        result, err := client.BRPop(ctx, 0*time.Second, "SUBMISSION_QUEUE").Result()
        if err != nil {
            fmt.Println("Error fetching job:", err)
            continue
        }

        var job Submission
        if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
            fmt.Println("Error decoding job:", err)
            continue
        }

        fmt.Printf("Evaluating submission ID: %s\n", job.ID)
        // Add your evaluation logic here
    }
}
