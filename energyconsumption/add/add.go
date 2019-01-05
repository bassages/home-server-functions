package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type SmartMeterMessage struct {
	Datumtijd                     string `json:"datumtijd"`
	StroomOpgenomenVermogenInWatt int    `json:"stroomOpgenomenVermogenInWatt"`
	StroomTariefIndicator         int    `json:"stroomTariefIndicator"`
}

type ElectricityUsage struct {
	Id              string `json:"id"`
	Timestamp       string `json:"datetime"`
	Watt            int    `json:"watt"`
	TariffIndicator int    `json:"tariffIndicator"`
}

func add(smartMeterMessage SmartMeterMessage) ElectricityUsage {
	id, uuidError := uuid.NewRandom()
	handleError(uuidError)

	electricityUsage := ElectricityUsage{Id: id.String(), Timestamp: smartMeterMessage.Datumtijd, Watt: smartMeterMessage.StroomOpgenomenVermogenInWatt, TariffIndicator: smartMeterMessage.StroomTariefIndicator}

	config := &aws.Config{
		Region: aws.String("eu-central-1"),
		// Use the default credentials chain, see https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials
		// "When you initialize a new service client without providing any credential arguments the SDK uses the default credential provider chain to find AWS credentials."
	}

	awsSession := session.Must(session.NewSession(config))
	dynamoDB := dynamodb.New(awsSession)

	itemToPut, marshallError := dynamodbattribute.MarshalMap(electricityUsage)
	handleError(marshallError)

	var putInput = &dynamodb.PutItemInput{
		TableName: aws.String("electricityusage"),
		Item:      itemToPut,
	}
	var _, putItemError = dynamoDB.PutItem(putInput)
	handleError(putItemError)

	return electricityUsage
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
