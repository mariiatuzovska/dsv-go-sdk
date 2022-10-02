package main

import (
	"fmt"
	"os"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	dsvclient "github.com/mariiatuzovska/dsv-go-sdk/go-swagger/client"
	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/client/secrets"
	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/client/tokens"
)

func main() {
	client := dsvclient.NewHTTPClientWithConfig(nil, dsvclient.DefaultTransportConfig().
		WithHost(os.Getenv("DSV_HOST")))
	username := os.Getenv("DSV_USERNAME")
	password := os.Getenv("DSV_PASSWORD")
	token, err := client.Tokens.Token(&tokens.TokenParams{
		GrantType: "password",
		Username:  &username,
		Password:  &password,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", token)
	secret, err := client.Secrets.GetSecret(&secrets.GetSecretParams{
		Path: os.Getenv("DSV_SECRET_PATH"),
	}, ClientCredentials(token.Payload.Token))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", secret)
}

func ClientCredentials(token string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("User-Agent", "go-cli")
		r.SetHeaderParam("Authorization", "Bearer "+token)
		return nil
	})
}
