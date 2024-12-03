package xpprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/provider"
)

type AzureClientBuilder clients.ClientBuilder

func GetProviderSchema(_ context.Context) (*schema.Provider, error) {
	return provider.AzureProvider(), nil
}

func (b *AzureClientBuilder) GetClient(ctx context.Context) (*clients.Client, error) {
	return clients.Build(ctx, (clients.ClientBuilder)(*b))
}
