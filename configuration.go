package apptree

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
	Name                 string         `json:"name"`
	Type                 Type           `json:"attributeType"`
	Index                int            `json:"attributeIndex"`
	RelatedConfiguration *Configuration `json:"relatedService,omitempty"`
}

type ConfigurationError struct {
	Message string
}

func (e ConfigurationError) Error() string {
	return e.Message
}
