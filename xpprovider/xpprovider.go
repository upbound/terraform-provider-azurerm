package xpprovider

import (
	"context"

	sdkclient "github.com/hashicorp/go-azure-sdk/sdk/client"
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

// RegisterResponseMiddleware appends mw to every go-azure-sdk client held by
// meta (the value returned by schema.Provider.Meta()). It returns false when
// meta is not a *clients.Client.
func RegisterResponseMiddleware(meta any, mw sdkclient.ResponseMiddleware) bool {
	c, ok := meta.(*clients.Client)
	if !ok {
		return false
	}
	c.AppendResponseMiddleware(mw)
	return true
}
