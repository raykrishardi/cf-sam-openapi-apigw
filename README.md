# cf-sam-openapi-apigw
Create Amazon API Gateway with OpenAPI spec with golambda and dynamodb (reference: https://github.com/aws-samples/cf-sam-openapi-file-organization-demo/tree/main)

## Getting Started

## Local Deployment

### Run on local machine
1. Deploy the stack
```
npm i -g cfn-include swagger-cli
make build
sam deploy --guided
```

2. Populate the DynamoDB with content based on this attribute definitions https://github.com/raykrishardi/cf-sam-openapi-apigw/blob/main/cloudformation/resources/dynamodb/animalsTable.yaml#L8-L12
  
   For example:
```
id: 1
name: Alligator
```

3. Hit the API GW
  
   For example:
```
curl https://<url>/Prod/animals/1
```

### Display OpenAPI yaml using swagger
```
docker run -d --name swagger-ui -p 8080:8080 -e SWAGGER_JSON=/api/api.yaml -v ./api:/api/ swaggerapi/swagger-ui
```

### Generate OpenAPI doc (response schema doesn't seem to get displayed, so use swagger UI instead)
```
npm i -g @openapitools/openapi-generator-cli
make generateApiDoc
open the apidocs/index.html file on your browser to view the documentation
```
