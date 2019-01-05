package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, smartMeterMessage SmartMeterMessage) (ElectricityUsage, error) {
	return add(smartMeterMessage), nil
}

func main() {
	lambda.Start(HandleRequest)
}
