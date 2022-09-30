package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	dsvclient "github.com/mariiatuzovska/dsv-go-sdk/swagger-codegen/client"
)

func main() {
	client := dsvclient.NewAPIClient(&dsvclient.Configuration{
		BasePath: os.Getenv("DSV_BASE_PATH"),
	})
	acsessTokenResponse, httpResp, err := client.TokensApi.Token(context.Background(), "password", &dsvclient.TokensApiTokenOpts{
		Username: optional.NewString(os.Getenv("DSV_USERNAME")),
		Password: optional.NewString(os.Getenv("DSV_PASSWORD")),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d %#v\n", httpResp.StatusCode, acsessTokenResponse)

	client = dsvclient.NewAPIClient(&dsvclient.Configuration{
		BasePath: os.Getenv("DSV_BASE_PATH"),
		DefaultHeader: map[string]string{
			"Authorization": acsessTokenResponse.AccessToken,
		},
	})
	secretResponse, httpResp, err := client.SecretsApi.GetSecret(context.Background(), os.Getenv("DSV_SECRET_PATH"), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d %#v\n", httpResp.StatusCode, secretResponse)
}
