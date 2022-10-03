# DSV Go SDK

The purpose of this application is to provide a simple service for storing and getting secrets.

## Generate swagger-codegen SDK

1. Install [swagger-codegen](https://github.com/swagger-api/swagger-codegen).
2. Generate models: `$ swagger-codegen generate -i ./path/to/swagger.json -l go -o ./swagger-codegen/models --model-name-prefix DSV -Dmodels`.
3. Generate API client: `$ swagger-codegen generate -i ./path/to/swagger.json -l go -o ./swagger-codegen/client --model-name-prefix DSV -c ./config.json`.

### Generate with docker

```
$ docker run --rm -v ${PWD}/swagger-codegen:/local swaggerapi/swagger-codegen-cli generate \
    -i /local/path/to/swagger.json \
    -l go \
    -o /local/client \
    --model-name-prefix DSV \
    -c /local/config.json
```

## Generate go-swagger

1. Install swagger: `$ go install github.com/go-swagger/go-swagger/cmd/swagger@latest`
2. Generate SDK: `$ swagger generate client -f ./path/to/swagger.json --client-package=./go-swagger/client --default-scheme=https --model-package=./go-swagger/models --api-package=./go-swagger/api`