package main

import (
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client"
)

type Config struct {
	Host  string
	TenantID string
}

func (c *Config) NewFMClient() (*client.Fleetmanager, error) {
	client := fleetmanager.NewHTTPClientWithHost(c.Host)

	return client, nil
}
