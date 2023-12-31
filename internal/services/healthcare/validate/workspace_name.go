// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"fmt"
	"regexp"
)

func WorkspaceName(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
		return
	}

	// The name attribute rules are :
	// 1. The value must be between 1 and 24 characters long
	// 2. must contain only lowercase letters or numbers.

	if !regexp.MustCompile(`^[a-z0-9]{1,24}$`).MatchString(v) {
		errors = append(errors, fmt.Errorf("%s must be between 1 and 24 characters long and can contain only lowercase letters or numbers", k))
		return
	}
	return warnings, errors
}
