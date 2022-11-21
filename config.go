package config

import (
	"errors"
	"io/ioutil"

	"github.com/Clinet/clinet_features"
	"github.com/JoshuaDoes/json"
)

type ConfigType int
const (
	ConfigTypeJSON ConfigType = iota
	ConfigTypeTOML
	ConfigTypeXML
)

type Config struct {
	Features     []features.Feature `json:"features"`

	path string //The path to the configuration file
}

//NewConfig returns a new configuration struct
func NewConfig() *Config {
	return &Config{
		Features: make([]features.Feature, 0),
	}
}

//LoadConfig creates a new configuration struct with the values in the specified configuration file
func LoadConfig(path string, cfgType ConfigType) (*Config, error) {
	cfg := &Config{path: path}

	switch cfgType {
	case ConfigTypeJSON:
		configJSON, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(configJSON, cfg)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("bot: config: unknown configuration type")
	}

	return cfg, nil
}

func SaveConfig(cfg *Config, path string, cfgType ConfigType) error {
	configJSON, err := json.Marshal(cfg, true)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, configJSON, 0644)
	return err
}

//LoadFrom loads the configuration from the specified path into the current cfg
func (cfg *Config) LoadFrom(path string, cfgType ConfigType) (err error) {
	cfg, err = LoadConfig(path, cfgType)
	return err
}
//SaveTo saves the current cfg to the specified path
func (cfg *Config) SaveTo(path string, cfgType ConfigType) (err error) {
	err = SaveConfig(cfg, path, cfgType)
	return err
}