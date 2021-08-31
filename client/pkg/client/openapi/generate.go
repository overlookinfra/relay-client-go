package openapi

//go:generate 	docker run --rm -v "${PWD}:/local" gcr.io/nebula-contrib/openapi-generator-cli:latest generate -i https://api.relay.sh/openapi/latest -g go --global-property apis,models,supportingFiles,modelDocs=false -o /local/pkg/client/openapi
