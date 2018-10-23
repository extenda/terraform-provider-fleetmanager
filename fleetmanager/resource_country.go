package fleetmanager

import (
	"log"
	"strings"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCountry() *schema.Resource {
	return &schema.Resource{
		Create: resourceCountryCreate,
		Read:   resourceCountryRead,
		Update: resourceCountryUpdate,
		Delete: resourceCountryDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"iso_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"brand_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCountryCreate(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	name := d.Get("name").(string)
	isoCode := d.Get("iso_code").(string)
	brandID := d.Get("brand_id").(string)

	params := fleet.NewCreateCountryParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithBody(&models.CreateFleetCountry{Name: &name, IsoCode: &isoCode})

	res, err := client.Fleet.CreateCountry(params)

	if err != nil {
		log.Printf("Failed to create country %q", err)

		return err
	}

	components := strings.Split(res.Location, "/")
	countryId := components[len(components)-1]
	d.SetId(countryId)

	return nil
}

func resourceCountryRead(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Id()

	params := fleet.NewGetCountryByIDParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)

	res, err := client.Fleet.GetCountryByID(params)

	if err != nil {
		log.Printf("Failed to find country %q", err)
		d.SetId("")

		return nil
	}

	d.Set("name", res.Payload.Name)
	d.Set("iso_code", res.Payload.IsoCode)

	return nil
}

func resourceCountryUpdate(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	name := d.Get("name").(string)
	brandID := d.Get("brand_id").(string)
	countryID := d.Id()

	params := fleet.NewUpdateCountryParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithBody(&models.UpdateFleetCountry{Name: &name})

	_, err := client.Fleet.UpdateCountry(params)

	if err != nil {
		log.Printf("Failed to find country %q", err)
		d.SetId("")

		return nil
	}

	return nil
}

func resourceCountryDelete(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Id()

	params := fleet.NewDeleteCountryParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)

	_, err := client.Fleet.DeleteCountry(params)

	if err != nil {
		log.Printf("Failed to find country %q", err)

		return nil
	}

	return nil
}
