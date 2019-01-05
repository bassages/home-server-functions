package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type ElectricityUsage struct {
	Id        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Watt      int    `json:"watt"`
}

func list() []ElectricityUsage {
	config := &aws.Config{
		Region: aws.String("eu-central-1"),
		// Use the default credentials chain, see https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials
		// "When you initialize a new service client without providing any credential arguments the SDK uses the default credential provider chain to find AWS credentials."
	}

	sess := session.Must(session.NewSession(config))
	svc := dynamodb.New(sess)

	var scanInput = &dynamodb.ScanInput{
		TableName: aws.String("electricityusage"),
	}
	var scanOutput, err = svc.Scan(scanInput)
	if err != nil {
		panic(err)
	}

	result := []ElectricityUsage{}
	err = dynamodbattribute.UnmarshalListOfMaps(scanOutput.Items, &result)
	if err != nil {
		panic(err)
	}

	id, _ := uuid.NewRandom()
	fmt.Println(id)

	return result

}
