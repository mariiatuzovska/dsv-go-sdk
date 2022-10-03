package main

import (
	"context"
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	dsvclient "github.com/mariiatuzovska/dsv-go-sdk/go-swagger/client"
	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/client/secrets"
	"github.com/mariiatuzovska/dsv-go-sdk/go-swagger/client/tokens"
)

func main() {
	transport := httptransport.New(os.Getenv("DSV_HOST"), "/v1", []string{"https"})
	client := dsvclient.New(transport, strfmt.Default)
	username := os.Getenv("DSV_USERNAME")
	password := os.Getenv("DSV_PASSWORD")
	token, err := client.Tokens.Token(&tokens.TokenParams{
		GrantType: "password",
		Username:  &username,
		Password:  &password,
		Context:   context.Background(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", token.Payload)
	secret, err := client.Secrets.GetSecret(&secrets.GetSecretParams{
		Path:    os.Getenv("DSV_SECRET_PATH"),
		Context: context.Background(),
	}, httptransport.BearerToken(token.Payload.Token))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", secret.Payload)
}
