package main

import (
	"log"
	"strings"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_brand"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
	"github.com/hashicorp/terraform/helper/schema"
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
	params.WithTenantID("hm")
	params.WithBody(&models.CreateFleetBrand{Name: &name})

	res, err := client.FleetBrand.PostFleetTenantTenantIDBrand(params)

	if err != nil {
		log.Printf("Failed to create brand %q", err)

		return err
	}

	components := strings.Split(res.Location, "/")
	brandId := components[len(components)-1]
	d.SetId(brandId)

	return nil
}

func resourceBrandRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Fleetmanager)

	params := fleet_brand.NewGetFleetTenantTenantIDBrandBrandIDParams()
	params.WithBrandID(d.Id())
	params.WithTenantID("hm")
	res, err := client.FleetBrand.GetFleetTenantTenantIDBrandBrandID(params)

	if err != nil {
		log.Printf("Failed to find brand %q", err)
		d.SetId("")

		return nil
	}

	d.Set("name", res.Payload.Name)

	return nil
}

func resourceBrandUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceBrandDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
