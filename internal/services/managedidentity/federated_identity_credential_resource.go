// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedidentity

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonschema"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedidentity/2023-01-31/managedidentities"
	"github.com/hashicorp/terraform-provider-azurerm/internal/locks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

var _ sdk.Resource = FederatedIdentityCredentialResource{}

const federatedIdentityCredentialPreviewAPIVersion = "2025-01-31-preview"

type FederatedIdentityCredentialResource struct{}

func (r FederatedIdentityCredentialResource) ModelObject() interface{} {
	return &FederatedIdentityCredentialResourceSchema{}
}

type FederatedIdentityCredentialResourceSchema struct {
	Audience                        []string `tfschema:"audience"`
	ClaimsMatchingExpressionValue   string   `tfschema:"claims_matching_expression_value"`
	ClaimsMatchingExpressionVersion int      `tfschema:"claims_matching_expression_version"`
	Issuer                          string   `tfschema:"issuer"`
	Name                            string   `tfschema:"name"`
	ResourceGroupName               string   `tfschema:"resource_group_name"`
	ResourceName                    string   `tfschema:"parent_id"`
	Subject                         string   `tfschema:"subject"`
}

func (r FederatedIdentityCredentialResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return managedidentities.ValidateFederatedIdentityCredentialID
}
func (r FederatedIdentityCredentialResource) ResourceType() string {
	return "azurerm_federated_identity_credential"
}
func (r FederatedIdentityCredentialResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"audience": {
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
			ForceNew: true,
			Required: true,
			Type:     pluginsdk.TypeList,
			MaxItems: 1,
		},
		"issuer": {
			ForceNew: true,
			Required: true,
			Type:     pluginsdk.TypeString,
		},
		"name": {
			ForceNew: true,
			Required: true,
			Type:     pluginsdk.TypeString,
		},
		"resource_group_name": commonschema.ResourceGroupName(),
		"parent_id": {
			// TODO: this wants renaming to `user_assigned_identity_id` (and `resource_group_name` removing in 4.0)
			Type:         pluginsdk.TypeString,
			ForceNew:     true,
			Required:     true,
			ValidateFunc: commonids.ValidateUserAssignedIdentityID,
		},
		"subject": {
			ForceNew: true,
			Optional: true,
			Type:     pluginsdk.TypeString,
			ExactlyOneOf: []string{
				"subject",
				"claims_matching_expression_value",
			},
			ConflictsWith: []string{
				"claims_matching_expression_value",
				"claims_matching_expression_version",
			},
		},
		"claims_matching_expression_value": {
			ForceNew: true,
			Optional: true,
			Type:     pluginsdk.TypeString,
			ExactlyOneOf: []string{
				"subject",
				"claims_matching_expression_value",
			},
			RequiredWith: []string{
				"claims_matching_expression_version",
			},
			ConflictsWith: []string{
				"subject",
			},
		},
		"claims_matching_expression_version": {
			ForceNew:     true,
			Optional:     true,
			Type:         pluginsdk.TypeInt,
			ValidateFunc: validation.IntAtLeast(1),
			RequiredWith: []string{
				"claims_matching_expression_value",
			},
			ConflictsWith: []string{
				"subject",
			},
		},
	}
}
func (r FederatedIdentityCredentialResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}
func (r FederatedIdentityCredentialResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ManagedIdentity.V20230131.ManagedIdentities
			baseUri, err := r.resourceManagerBaseURI(metadata)
			if err != nil {
				return err
			}

			var config FederatedIdentityCredentialResourceSchema
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			subscriptionId := metadata.Client.Account.SubscriptionId
			parentId, err := commonids.ParseUserAssignedIdentityID(config.ResourceName)
			if err != nil {
				return fmt.Errorf("parsing parent resource ID: %+v", err)
			}

			locks.ByID(parentId.ID())
			defer locks.UnlockByID(parentId.ID())

			id := managedidentities.NewFederatedIdentityCredentialID(subscriptionId, config.ResourceGroupName, parentId.UserAssignedIdentityName, config.Name)

			existing, err := r.getFederatedIdentityCredential(ctx, client, baseUri, id)
			if err != nil {
				if !response.WasNotFound(existing.HttpResponse) {
					return fmt.Errorf("checking for the presence of an existing %s: %+v", id, err)
				}
			}
			if !response.WasNotFound(existing.HttpResponse) {
				return metadata.ResourceRequiresImport(r.ResourceType(), id)
			}

			payload := r.mapFederatedIdentityCredentialResourceSchemaToPreviewPayload(config)
			if _, err := r.createOrUpdateFederatedIdentityCredential(ctx, client, baseUri, id, payload); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}
func (r FederatedIdentityCredentialResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ManagedIdentity.V20230131.ManagedIdentities
			baseUri, err := r.resourceManagerBaseURI(metadata)
			if err != nil {
				return err
			}
			schema := FederatedIdentityCredentialResourceSchema{}

			id, err := managedidentities.ParseFederatedIdentityCredentialID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := r.getFederatedIdentityCredential(ctx, client, baseUri, *id)
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(*id)
				}
				return fmt.Errorf("retrieving %s: %+v", *id, err)
			}

			if model := resp.Model; model != nil {
				schema.Name = id.FederatedIdentityCredentialName
				schema.ResourceGroupName = id.ResourceGroupName
				parentId := commonids.NewUserAssignedIdentityID(id.SubscriptionId, id.ResourceGroupName, id.UserAssignedIdentityName)
				schema.ResourceName = parentId.ID()
				r.mapPreviewPayloadToFederatedIdentityCredentialResourceSchema(*model, &schema)
			}

			return metadata.Encode(&schema)
		},
	}
}
func (r FederatedIdentityCredentialResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 30 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.ManagedIdentity.V20230131.ManagedIdentities
			baseUri, err := r.resourceManagerBaseURI(metadata)
			if err != nil {
				return err
			}

			var config FederatedIdentityCredentialResourceSchema
			if err := metadata.Decode(&config); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			parentId, err := commonids.ParseUserAssignedIdentityID(config.ResourceName)
			if err != nil {
				return fmt.Errorf("parsing parent resource ID: %+v", err)
			}

			locks.ByID(parentId.ID())
			defer locks.UnlockByID(parentId.ID())

			id, err := managedidentities.ParseFederatedIdentityCredentialID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			if _, err := r.deleteFederatedIdentityCredential(ctx, client, baseUri, *id); err != nil {
				return fmt.Errorf("deleting %s: %+v", *id, err)
			}

			return nil
		},
	}
}

type FederatedIdentityCredentialClaimsMatchingExpression struct {
	LanguageVersion int64  `json:"languageVersion"`
	Value           string `json:"value"`
}

type FederatedIdentityCredentialPreviewProperties struct {
	Audiences                []string                                             `json:"audiences"`
	ClaimsMatchingExpression *FederatedIdentityCredentialClaimsMatchingExpression `json:"claimsMatchingExpression,omitempty"`
	Issuer                   string                                               `json:"issuer"`
	Subject                  *string                                              `json:"subject,omitempty"`
}

type FederatedIdentityCredentialPreviewPayload struct {
	Properties *FederatedIdentityCredentialPreviewProperties `json:"properties,omitempty"`
}

func (r FederatedIdentityCredentialResource) mapFederatedIdentityCredentialResourceSchemaToPreviewPayload(input FederatedIdentityCredentialResourceSchema) FederatedIdentityCredentialPreviewPayload {
	properties := &FederatedIdentityCredentialPreviewProperties{
		Audiences: input.Audience,
		Issuer:    input.Issuer,
	}
	if input.Subject != "" {
		subject := input.Subject
		properties.Subject = &subject
	}
	if input.ClaimsMatchingExpressionValue != "" {
		properties.ClaimsMatchingExpression = &FederatedIdentityCredentialClaimsMatchingExpression{
			LanguageVersion: int64(input.ClaimsMatchingExpressionVersion),
			Value:           input.ClaimsMatchingExpressionValue,
		}
	}
	return FederatedIdentityCredentialPreviewPayload{Properties: properties}
}

func (r FederatedIdentityCredentialResource) mapPreviewPayloadToFederatedIdentityCredentialResourceSchema(input FederatedIdentityCredentialPreviewPayload, output *FederatedIdentityCredentialResourceSchema) {
	if input.Properties == nil {
		input.Properties = &FederatedIdentityCredentialPreviewProperties{}
	}
	output.Audience = input.Properties.Audiences
	output.Issuer = input.Properties.Issuer
	if input.Properties.Subject != nil {
		output.Subject = *input.Properties.Subject
	}
	if input.Properties.ClaimsMatchingExpression != nil {
		output.ClaimsMatchingExpressionValue = input.Properties.ClaimsMatchingExpression.Value
		output.ClaimsMatchingExpressionVersion = int(input.Properties.ClaimsMatchingExpression.LanguageVersion)
	}
}

func (r FederatedIdentityCredentialResource) resourceManagerBaseURI(metadata sdk.ResourceMetaData) (string, error) {
	endpoint, ok := metadata.Client.Account.Environment.ResourceManager.Endpoint()
	if !ok {
		return "", fmt.Errorf("resolving resource manager endpoint for preview federated identity credential API")
	}
	return *endpoint, nil
}

type federatedIdentityCredentialsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *FederatedIdentityCredentialPreviewPayload
}

type federatedIdentityCredentialsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

func (r FederatedIdentityCredentialResource) createOrUpdateFederatedIdentityCredential(ctx context.Context, c *managedidentities.ManagedIdentitiesClient, baseUri string, id managedidentities.FederatedIdentityCredentialId, payload FederatedIdentityCredentialPreviewPayload) (*http.Response, error) {
	queryParameters := map[string]interface{}{
		"api-version": federatedIdentityCredentialPreviewAPIVersion,
	}
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(payload),
		autorest.WithQueryParameters(queryParameters),
	)
	req, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return resp, err
	}
	return resp, autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK), autorest.ByClosing())
}

func (r FederatedIdentityCredentialResource) getFederatedIdentityCredential(ctx context.Context, c *managedidentities.ManagedIdentitiesClient, baseUri string, id managedidentities.FederatedIdentityCredentialId) (result federatedIdentityCredentialsGetOperationResponse, err error) {
	queryParameters := map[string]interface{}{
		"api-version": federatedIdentityCredentialPreviewAPIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters),
	)
	req, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedidentity.FederatedIdentityCredentialResource", "getFederatedIdentityCredential", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedidentity.FederatedIdentityCredentialResource", "getFederatedIdentityCredential", result.HttpResponse, "Failure sending request")
		return
	}

	err = autorest.Respond(
		result.HttpResponse,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing(),
	)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedidentity.FederatedIdentityCredentialResource", "getFederatedIdentityCredential", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

func (r FederatedIdentityCredentialResource) deleteFederatedIdentityCredential(ctx context.Context, c *managedidentities.ManagedIdentitiesClient, baseUri string, id managedidentities.FederatedIdentityCredentialId) (result federatedIdentityCredentialsDeleteOperationResponse, err error) {
	queryParameters := map[string]interface{}{
		"api-version": federatedIdentityCredentialPreviewAPIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters),
	)
	req, err := preparer.Prepare((&http.Request{}).WithContext(ctx))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedidentity.FederatedIdentityCredentialResource", "deleteFederatedIdentityCredential", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedidentity.FederatedIdentityCredentialResource", "deleteFederatedIdentityCredential", result.HttpResponse, "Failure sending request")
		return
	}

	err = autorest.Respond(result.HttpResponse, azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK), autorest.ByClosing())
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedidentity.FederatedIdentityCredentialResource", "deleteFederatedIdentityCredential", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}
