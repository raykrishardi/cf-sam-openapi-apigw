package main

import (
	"cf-sam-openapi-apigw/internal/entity"
	"cf-sam-openapi-apigw/internal/pkg/utils"
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"cf-sam-openapi-apigw/internal/pkg/config"
	ddbrepo "cf-sam-openapi-apigw/internal/pkg/dynamodb"
	ddbuc "cf-sam-openapi-apigw/internal/usecase/dynamodb"
)

var (
	AWS_REGION = os.Getenv("AWS_REGION")
	TABLE_NAME = os.Getenv("TABLE_NAME")
)

func handler(ctx context.Context, event entity.CustomAPIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	utils.InfoLog.Printf("event: %+v\n", event)

	// Initialise app config
	appConfig := &config.AppConfig{
		AWSRegion: AWS_REGION,
		TableName: TABLE_NAME,
	}

	// Initialise repositories
	ddbRepo := ddbrepo.NewDynamodbRepository(appConfig)
	ddbrepo.NewDynamodb(ddbRepo)

	// Initialise specific usecases
	ddbUC := ddbuc.NewDynamodbUseCase(ctx, ddbRepo)

	// Get animal by ID
	getAnimalInput := ddbuc.GetAnimalInput{
		TableName: appConfig.TableName,
		AnimalID:  event.AnimalID,
	}
	getAnimalOutput, err := ddbUC.GetAnimalByID(ctx, getAnimalInput)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.New("error: unable to get animal by ID")
	}
	utils.InfoLog.Printf("getAnimalOutput: %+v\n", getAnimalOutput)

	getAnimalOutputJSON, err := json.Marshal(getAnimalOutput)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.New("error: unable to marshal getAnimalOutput")
	}

	return events.APIGatewayProxyResponse{Body: string(getAnimalOutputJSON), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
