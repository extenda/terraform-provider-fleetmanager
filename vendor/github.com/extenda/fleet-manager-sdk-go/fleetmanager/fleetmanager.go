package fleetmanager

import (
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client"
	"github.com/go-openapi/strfmt"
)

func NewHTTPClientWithHost(host string) *client.Fleetmanager {
	basePath := "/v1"
	t := client.TransportConfig{Host: host,
		BasePath: basePath, Schemes: client.DefaultSchemes}

	return client.NewHTTPClientWithConfig(strfmt.Default, &t)
}
