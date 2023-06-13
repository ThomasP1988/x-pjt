package main

import (
	"NFTM/shared/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	entity "NFTM/shared/entities/notification"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
	config.GetConfig(nil)
	lambda.Start(func(ctx context.Context, req DynamoDBEvent) {
		for _, record := range req.Records {
			if events.DynamoDBOperationType(record.EventName) == events.DynamoDBOperationTypeInsert {
				fmt.Printf("record.Change.NewImage: %+v\n", record.Change.NewImage)
				notification := &entity.Notification{}
				err := dynamodbattribute.UnmarshalMap(record.Change.NewImage, &notification)

				if err != nil {
					fmt.Printf("err: %v\n", err)
					continue
				}

				jsonData := map[string]interface{}{
					"operationName": "notify",
					"query": `mutation notify($notification: NotificationInput) {
							notify(input: $notification) {
								id,
							}
						}`,
					"variables": map[string]interface{}{
						"notification": notification,
					},
				}
				jsonValue, err := json.Marshal(jsonData)
				if err != nil {
					fmt.Printf("Error marshalling json request %s\n", err)
				}
				fmt.Printf("jsonValue: %v\n", string(jsonValue))
				request, err := http.NewRequest("POST", os.Getenv("APIURL"), bytes.NewBuffer(jsonValue))
				request.Header.Add("x-api-key", os.Getenv("APIKey"))
				if err != nil {
					fmt.Printf("Error setting request %s\n", err)
				}

				client := &http.Client{Timeout: time.Second * 10}
				response, err := client.Do(request)
				if err != nil {
					fmt.Printf("The HTTP request failed with error %s\n", err)
				}

				defer response.Body.Close()
				data, _ := ioutil.ReadAll(response.Body)
				fmt.Printf("data: %v\n", data)

				fmt.Printf("string(data): %v\n", string(data))
			}
		}
	})
}

type DynamoDBEvent struct {
	Records []DynamoDBEventRecord `json:"Records"`
}

type DynamoDBEventRecord struct {
	AWSRegion      string                       `json:"awsRegion"`
	Change         DynamoDBStreamRecord         `json:"dynamodb"`
	EventID        string                       `json:"eventID"`
	EventName      string                       `json:"eventName"`
	EventSource    string                       `json:"eventSource"`
	EventVersion   string                       `json:"eventVersion"`
	EventSourceArn string                       `json:"eventSourceARN"`
	UserIdentity   *events.DynamoDBUserIdentity `json:"userIdentity,omitempty"`
}

type DynamoDBStreamRecord struct {
	ApproximateCreationDateTime events.SecondsEpochTime `json:"ApproximateCreationDateTime,omitempty"`
	// changed to map[string]*dynamodb.AttributeValue
	Keys map[string]*dynamodb.AttributeValue `json:"Keys,omitempty"`
	// changed to map[string]*dynamodb.AttributeValue
	NewImage map[string]*dynamodb.AttributeValue `json:"NewImage,omitempty"`
	// changed to map[string]*dynamodb.AttributeValue
	OldImage       map[string]*dynamodb.AttributeValue `json:"OldImage,omitempty"`
	SequenceNumber string                              `json:"SequenceNumber"`
	SizeBytes      int64                               `json:"SizeBytes"`
	StreamViewType string                              `json:"StreamViewType"`
}
