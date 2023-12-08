package dynamodb

import (
	"cf-sam-openapi-apigw/internal/entity"
	ddbrepo "cf-sam-openapi-apigw/internal/pkg/dynamodb"
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamodbUseCase struct {
	DynamodbRepo *ddbrepo.DynamodbRepository
	Client       *dynamodb.Client
}

func NewDynamodbUseCase(ctx context.Context, ddbRepo *ddbrepo.DynamodbRepository) *DynamodbUseCase {
	uc := &DynamodbUseCase{}

	client, err := ddbRepo.GetDynamodbClient(ctx)
	if err != nil {
		return uc
	}

	uc.DynamodbRepo = ddbRepo
	uc.Client = client

	return uc
}

func (uc *DynamodbUseCase) GetAnimalByID(ctx context.Context, params GetAnimalInput) (entity.Animal, error) {
	a := entity.Animal{}

	gii := &dynamodb.GetItemInput{
		TableName: &params.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: params.AnimalID},
		},
	}

	gio, err := uc.Client.GetItem(ctx, gii)
	if err != nil {
		return a, err
	}

	if len(gio.Item) < 1 {
		return a, errors.New("animal with that ID is not found")
	}

	attributevalue.UnmarshalMap(gio.Item, &a)

	return a, nil
}
