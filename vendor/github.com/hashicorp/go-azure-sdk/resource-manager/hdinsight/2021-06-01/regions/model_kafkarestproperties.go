package regions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KafkaRestProperties struct {
	ClientGroupInfo       *ClientGroupInfo   `json:"clientGroupInfo,omitempty"`
	ConfigurationOverride *map[string]string `json:"configurationOverride,omitempty"`
}
