package model

import "encoding/json"

// ObjectConfig config of objects for code generation
type ObjectConfig struct {
	Objects []SolarisObject `json:"Objects"`
}

// NewSolarisConfig creates a new config instance
func NewSolarisConfig() *ObjectConfig {
	return &ObjectConfig{
		Objects: []SolarisObject{},
	}
}

// MarshalJson serializes the config as json
func (config *ObjectConfig) MarshalJson() ([]byte, error) {
	return json.Marshal(config)
}

// UnmarshalJson retrieves the config from json data
func (config *ObjectConfig) UnmarshalJson(data []byte) error {
	if err := json.Unmarshal(data, config); err != nil {
		return err
	}
	return nil
}
