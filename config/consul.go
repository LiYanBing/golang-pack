package config

import (
	"errors"

	"github.com/hashicorp/consul/api"
)

type Consul struct {
	Address    string `json:"address,omitempty"`     // address of config consul
	DataCenter string `json:"data_center,omitempty"` // data center of consul
	ConfigPath string `json:"config_path"`           // project config path in the consul
}

var EmptyPathError = errors.New("empty config_path")

// get project config from consul
func GetConfigFromConsul(consul Consul) ([]byte, error) {
	if consul.ConfigPath == "" {
		return nil, EmptyPathError
	}

	client, err := api.NewClient(&api.Config{
		Address:    consul.Address,
		Datacenter: consul.DataCenter,
	})
	if err != nil {
		return nil, err
	}

	pair, _, err := client.KV().Get(consul.ConfigPath, nil)
	if err != nil {
		return nil, err
	}

	return pair.Value, nil
}
