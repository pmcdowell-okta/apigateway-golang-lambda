package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	datas := make(map[string]string)

	datas["Dude where is my"]=request.Body

	returnString, err := json.Marshal(datas)
	_=err

	return events.APIGatewayProxyResponse{Body: string(returnString), StatusCode: 200}, nil
}

func main() {

	lambda.Start(Handler)
}