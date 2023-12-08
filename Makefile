# Compile openAPI components into a single file
buildOpenApi:
	swagger-cli bundle -o api/api.yaml --dereference --t yaml api/apiSkeleton.yaml
	swagger-cli validate api/api.yaml

build: buildOpenApi
	cfn-include --yaml cloudFormation/templateSkeleton.yaml > template.yaml
	cd golambda && go mod download
	rm -rf .aws-sam
	sam build -t template.yaml

deploy:
	sam deploy --no-confirm-changeset