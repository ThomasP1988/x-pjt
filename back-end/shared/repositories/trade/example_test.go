package trade

import (
	"context"
	"fmt"

	"github.com/influxdata/influxdb-client-go/v2/api"
)

func ExampleDeleteBuckets() {
	client, org, err := getClientAndOrganisation()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	buckets, err := (*client).BucketsAPI().GetBuckets(context.Background())

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	for _, bucket := range *buckets {
		err := (*client).BucketsAPI().DeleteBucket(context.Background(), &bucket)

		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		fmt.Printf("deleted bucket.Name: %v\n", bucket.Name)
	}

	tasks, err := (*client).TasksAPI().FindTasks(context.Background(), &api.TaskFilter{
		OrgID: *org.Id,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	for _, task := range tasks {
		err := (*client).TasksAPI().DeleteTask(context.Background(), &task)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		fmt.Printf("deleted task.Name: %v\n", task.Name)
	}
	// output: yes
}
