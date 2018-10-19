package fleetmanager

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FM_HOST", nil),
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FM_TENANT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fleetmanager_brand": resourceBrand(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Host:     d.Get("host").(string),
		TenantID: d.Get("tenant_id").(string),
	}

	return config.NewFMClient()
}
