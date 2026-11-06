package main

import (
	"context"
	"fmt"
	"log"

	"github.com/capturekitdev/capturekit-go"
)

func main() {
	client, err := capturekit.NewClient("", nil) // reads CAPTUREKIT_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	submitted, err := client.Content(ctx, &capturekit.ContentParams{URL: "https://example.com", IsAsync: capturekit.Bool(true)})
	if err != nil {
		log.Fatal(err)
	}
	var accepted struct {
		JobID string `json:"job_id"`
	}
	if err := submitted.Decode(&accepted); err != nil {
		log.Fatal(err)
	}
	job, err := client.WaitForJob(ctx, accepted.JobID, &capturekit.WaitForJobOptions{
		OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(job.Body))
}
