---
subcategory: "Authorization"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_federated_identity_credential"
description: |-
  Manages a Federated Identity Credential.
---

# azurerm_federated_identity_credential

Manages a Federated Identity Credential.

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "example"
  location = "West Europe"
}

resource "azurerm_user_assigned_identity" "example" {
  location            = azurerm_resource_group.example.location
  name                = "example"
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_federated_identity_credential" "example" {
  name                = "example"
  resource_group_name = azurerm_resource_group.example.name
  audience            = ["foo"]
  issuer              = "https://foo"
  parent_id           = azurerm_user_assigned_identity.example.id
  subject             = "foo"
}
```

## Example Usage (claims matching expression - preview API)

```hcl
resource "azurerm_resource_group" "example" {
  name     = "example"
  location = "West Europe"
}

resource "azurerm_user_assigned_identity" "example" {
  location            = azurerm_resource_group.example.location
  name                = "example"
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_federated_identity_credential" "example" {
  name                               = "example"
  resource_group_name                = azurerm_resource_group.example.name
  audience                           = ["api://AzureADTokenExchange"]
  issuer                             = "https://oidc.prod-aks.azure.com/<tenant>/<issuer>"
  parent_id                          = azurerm_user_assigned_identity.example.id
  claims_matching_expression_value   = "claims['sub'] matches 'system:serviceaccount:example-*:*'"
  claims_matching_expression_version = 1
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) Specifies the name of this Federated Identity Credential. Changing this forces a new Federated Identity Credential to be created.

* `resource_group_name` - (Required) Specifies the name of the Resource Group within which this Federated Identity Credential should exist. Changing this forces a new Federated Identity Credential to be created.

* `audience` - (Required) Specifies the audience for this Federated Identity Credential. Changing this forces a new Federated Identity Credential to be created.

* `issuer` - (Required) Specifies the issuer of this Federated Identity Credential. Changing this forces a new Federated Identity Credential to be created.

* `parent_id` - (Required) Specifies parent ID of User Assigned Identity for this Federated Identity Credential. Changing this forces a new Federated Identity Credential to be created.

* `subject` - (Optional) Specifies the subject for this Federated Identity Credential. Changing this forces a new Federated Identity Credential to be created.

* `claims_matching_expression_value` - (Optional) Specifies the claims matching expression value for this Federated Identity Credential. This uses the preview API behavior. Changing this forces a new Federated Identity Credential to be created.

* `claims_matching_expression_version` - (Optional) Specifies the claims matching expression language version for this Federated Identity Credential. This uses the preview API behavior and must be provided together with `claims_matching_expression_value`. Changing this forces a new Federated Identity Credential to be created.

~> **Note:** Exactly one of `subject` or `claims_matching_expression_value` must be specified.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the Federated Identity Credential.

---

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating this Federated Identity Credential.
* `delete` - (Defaults to 30 minutes) Used when deleting this Federated Identity Credential.
* `read` - (Defaults to 5 minutes) Used when retrieving this Federated Identity Credential.
* `update` - (Defaults to 30 minutes) Used when updating this Federated Identity Credential.

## Import

An existing Federated Identity Credential can be imported into Terraform using the `resource id`, e.g.

```shell
terraform import azurerm_federated_identity_credential.example /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{parentIdentityName}/federatedIdentityCredentials/{resourceName}
```
