// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"regexp"

	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

func RecoveryServicesVaultName(v interface{}, k string) (warnings []string, errors []error) {
	regexpValidator := validation.StringMatch(
		regexp.MustCompile("^[a-zA-Z][-a-zA-Z0-9]{1,49}$"),
		"Recovery Service Vault name must be 2 - 50 characters long, start with a letter, contain only letters, numbers and hyphens.",
	)
	return regexpValidator(v, k)
}
