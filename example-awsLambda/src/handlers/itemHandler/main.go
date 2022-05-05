package main

import (
	plog "example-awsLambda/src/customlog"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/valyala/fastjson"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	plog.PrintJson("invoked : ", request)
	ApiResponse := events.APIGatewayProxyResponse{}
	switch request.HTTPMethod {
	case "GET":
		username := request.QueryStringParameters["username"]
		if username != "" {
			ApiResponse = events.APIGatewayProxyResponse{Body: "Found username : " + username, StatusCode: 200}
		} else {
			ApiResponse = events.APIGatewayProxyResponse{Body: "Error: Query Parameter username missing", StatusCode: 500}
		}

	case "POST":
		err := fastjson.Validate(request.Body)

		if err != nil {
			body := "Error: Invalid JSON payload ||| " + fmt.Sprint(err) + " Body Obtained" + "||||" + request.Body
			ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 500}
		} else {
			ApiResponse = events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}
		}

	}
	return ApiResponse, nil
}

func main() {
	plog.PrintError("userHandler starting...\n")
	lambda.Start(HandleRequest)
}
