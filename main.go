package main

import (
	"github.com/hashicorp/terraform/plugin"

	fm "github.com/extenda/terraform-provider-fleetmanager/fleetmanager"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fm.Provider})
}
