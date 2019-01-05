package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) ([]ElectricityUsage, error) {
	return list(), nil
}

func main() {
	lambda.Start(HandleRequest)
}
