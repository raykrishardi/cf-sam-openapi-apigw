package dynamodb

import "cf-sam-openapi-apigw/internal/pkg/config"

var Repo *DynamodbRepository

type DynamodbRepository struct {
	App *config.AppConfig
}

func NewDynamodbRepository(app *config.AppConfig) *DynamodbRepository {
	return &DynamodbRepository{
		App: app,
	}
}

func NewDynamodb(repo *DynamodbRepository) {
	Repo = repo
}
