# Compile openAPI components into a single file
buildOpenApi:
	swagger-cli bundle -o api/api.yaml --dereference --t yaml api/apiSkeleton.yaml
	swagger-cli validate api/api.yaml

# Generating the OpenApi documentation
generateApiDoc:
	make buildOpenApi
	openapi-generator generate -i api/api.yaml -g html2 -o ./apidocs

build:
	make buildOpenApi
	cfn-include --yaml cloudFormation/templateSkeleton.yaml > template.yaml
	cd golambda && go mod download
	rm -rf .aws-sam
	sam build -t template.yaml

deploy:
	sam deploy --no-confirm-changeset