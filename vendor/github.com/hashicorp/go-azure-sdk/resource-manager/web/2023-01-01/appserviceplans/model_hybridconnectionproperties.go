package appserviceplans

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridConnectionProperties struct {
	Hostname            *string `json:"hostname,omitempty"`
	Port                *int64  `json:"port,omitempty"`
	RelayArmUri         *string `json:"relayArmUri,omitempty"`
	RelayName           *string `json:"relayName,omitempty"`
	SendKeyName         *string `json:"sendKeyName,omitempty"`
	SendKeyValue        *string `json:"sendKeyValue,omitempty"`
	ServiceBusNamespace *string `json:"serviceBusNamespace,omitempty"`
	ServiceBusSuffix    *string `json:"serviceBusSuffix,omitempty"`
}
