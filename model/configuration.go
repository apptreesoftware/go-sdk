package model

import (
	"errors"
	"fmt"
)

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

func (config Configuration) getConfigurationAttribute(index int) (*ConfigurationAttribute, error) {
	for _, element := range config.Attributes {
		if element.Index == index {
			return &element, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("No element found with index %d", index))
}

type ConfigurationAttribute struct {
	Name                 string        `json:"name"`
	Type                 Type          `json:"attributeType"`
	Index                int           `json:"attributeIndex"`
	RelatedConfiguration Configuration `json:"relatedService"`
}

type ConfigurationError struct {
	Message string
}

func (e ConfigurationError) Error() string {
	return e.Message
}
