package encodings

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Codec = DDAudio{}

type DDAudio struct {
	Bitrate      *int64 `json:"bitrate,omitempty"`
	Channels     *int64 `json:"channels,omitempty"`
	SamplingRate *int64 `json:"samplingRate,omitempty"`

	// Fields inherited from Codec
	Label *string `json:"label,omitempty"`
}

var _ json.Marshaler = DDAudio{}

func (s DDAudio) MarshalJSON() ([]byte, error) {
	type wrapper DDAudio
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DDAudio: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DDAudio: %+v", err)
	}
	decoded["@odata.type"] = "#Microsoft.Media.DDAudio"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DDAudio: %+v", err)
	}

	return encoded, nil
}
