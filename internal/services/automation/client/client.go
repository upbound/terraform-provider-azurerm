// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/webhook"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/agentregistrationinformation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/softwareupdateconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/watcher"
	automation_2023_11_01 "github.com/hashicorp/go-azure-sdk/resource-manager/automation/2023-11-01"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	*automation_2023_11_01.Client

	AgentRegistrationInfoClient *agentregistrationinformation.AgentRegistrationInformationClient
	SoftwareUpdateConfigClient  *softwareupdateconfiguration.SoftwareUpdateConfigurationClient
	WebhookClient               *webhook.WebhookClient
	WatcherClient               *watcher.WatcherClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	metaClient, err := automation_2023_11_01.NewClientWithBaseURI(o.Environment.ResourceManager, func(c *resourcemanager.Client) {
		o.Configure(c, o.Authorizers.ResourceManager)
	})
	if err != nil {
		return nil, fmt.Errorf("building Automation client: %+v", err)
	}

	agentRegistrationInfoClient := agentregistrationinformation.NewAgentRegistrationInformationClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&agentRegistrationInfoClient.Client, o.ResourceManagerAuthorizer)

	softUpClient := softwareupdateconfiguration.NewSoftwareUpdateConfigurationClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&softUpClient.Client, o.ResourceManagerAuthorizer)

	watcherClient := watcher.NewWatcherClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&watcherClient.Client, o.ResourceManagerAuthorizer)

	webhookClient := webhook.NewWebhookClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&webhookClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		Client: metaClient,

		AgentRegistrationInfoClient: &agentRegistrationInfoClient,
		SoftwareUpdateConfigClient:  &softUpClient,
		WatcherClient:               &watcherClient,
		WebhookClient:               &webhookClient,
	}, nil
}
