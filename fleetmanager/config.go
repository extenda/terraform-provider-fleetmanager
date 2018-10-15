package fleetmanager

import (
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client"
)

type Config struct {
	Host     string
	TenantID string
}

func (c *Config) NewFMClient() *client.Fleetmanager {
	client := fleetmanager.NewHTTPClientWithHost(c.Host)

	return client
}

type Tenant struct {
	ID     string
	Client *client.Fleetmanager
}

func (c *Config) Client() (interface{}, error) {
	var tenant Tenant
	tenant.ID = c.TenantID
	tenant.Client = c.NewFMClient()

	return &tenant, nil
}
