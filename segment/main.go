package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gopkg.in/segmentio/analytics-go.v3"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

type Data struct {
	EventName string
	UserId    string
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(sqsEvent events.SQSEvent) (Response, error) {
	segmentToken, ok := os.LookupEnv("SEGMENT_TOKEN")
	statusCode := 200

	if ok {
		client := analytics.New(segmentToken)

		var payload Data

		for _, message := range sqsEvent.Records {

			fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)

			err := json.Unmarshal([]byte(message.Body), &payload)

			if err != nil {
				fmt.Println("error:", err)
			} else {

				fmt.Printf("payload: %+v\n", payload)

				client.Enqueue(analytics.Track{
					Event:  payload.EventName,
					UserId: payload.UserId,
				})
			}
		}

		defer client.Close()
	}

	resp := Response{
		StatusCode: statusCode,
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
