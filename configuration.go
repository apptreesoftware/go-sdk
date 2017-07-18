package apptree

import "encoding/json"

type Configuration struct {
	Name       string                   `json:"name"`
	Attributes []ConfigurationAttribute `json:"attributes"`
}

func (config Configuration) MaxAttributeIndex() int {
	maxIndex := 0
	for _, attr := range config.Attributes {
		if maxIndex < attr.Index {
			maxIndex = attr.Index
		}
	}
	return maxIndex
}

func (config Configuration) getConfigurationAttribute(index int) *ConfigurationAttribute {
	for _, element := range config.Attributes {
		if element.Index == index {
			return &element
		}
	}
	return nil
}

type ConfigurationAttribute struct {
	Name                     string         `json:"name"`
	Type                     Type           `json:"attributeType"`
	Index                    int            `json:"attributeIndex"`
	RelatedConfiguration     *Configuration `json:"relatedService,omitempty"`
	RelatedListConfiguration *Configuration `json:"relatedListServiceConfiguration,omitempty"`
}

type configAttributeHelper struct {
	Name                     string         `json:"name"`
	Label                    string         `json:"label"`
	Type                     Type           `json:"attributeType"`
	Index                    int            `json:"attributeIndex"`
	RelatedConfiguration     *Configuration `json:"relatedService,omitempty"`
	RelatedListConfiguration *Configuration `json:"relatedListServiceConfiguration"`
}

func (a *ConfigurationAttribute) UnmarshalJSON(data []byte) error {
	var helper configAttributeHelper
	err := json.Unmarshal(data, &helper)
	if err != nil {
		return err
	}
	a.Name = helper.Name
	if helper.Name == "" {
		a.Name = helper.Label
	}
	a.Type = helper.Type
	a.Index = helper.Index
	a.RelatedConfiguration = helper.RelatedConfiguration
	a.RelatedListConfiguration = helper.RelatedListConfiguration
	return nil
}

type ConfigurationError struct {
	Message string
}

func (e ConfigurationError) Error() string {
	return e.Message
}
