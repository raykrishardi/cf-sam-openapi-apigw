package entity

type CustomAPIGatewayProxyRequest struct {
	Path     string `json:"path"`
	UserARN  string `json:"userArn"`
	AnimalID string `json:"animalId"`
}
