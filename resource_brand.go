package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_brand"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

func resourceBrand() *schema.Resource {
	return &schema.Resource{
		Create: resourceBrandCreate,
		Read:   resourceBrandRead,
		Update: resourceBrandUpdate,
		Delete: resourceBrandDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceBrandCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	client := m.(*client.Fleetmanager)

	params := fleet_brand.NewPostFleetTenantTenantIDBrandParams()
	params.WithTenantID("HM")
	params.WithBody(&models.CreateFleetBrand{Name: &name})

	res, err := client.FleetBrand.PostFleetTenantTenantIDBrand(params)

	if err != nil {
		log.Printf("Failed to creare brand %q", err)
		return err
	}

	log.Printf(res.Location)

	return nil
}

func resourceBrandRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceBrandUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceBrandDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
