package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsWorkFromAnywhereMetric{}

type UserExperienceAnalyticsWorkFromAnywhereMetric struct {
	// The work from anywhere metric devices. Read-only.
	MetricDevices *[]UserExperienceAnalyticsWorkFromAnywhereDevice `json:"metricDevices,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserExperienceAnalyticsWorkFromAnywhereMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsWorkFromAnywhereMetric{}

func (s UserExperienceAnalyticsWorkFromAnywhereMetric) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsWorkFromAnywhereMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsWorkFromAnywhereMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsWorkFromAnywhereMetric: %+v", err)
	}

	delete(decoded, "metricDevices")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsWorkFromAnywhereMetric: %+v", err)
	}

	return encoded, nil
}