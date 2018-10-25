package fleetmanager

import (
	"log"
	"strings"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceWorkstation() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkstationCreate,
		Read:   resourceWorkstationRead,
		Update: resourceWorkstationUpdate,
		Delete: resourceWorkstationDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"brand_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"country_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"store_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceWorkstationCreate(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Get("country_id").(string)
	storeID := d.Get("store_id").(string)
	name := d.Get("name").(string)

	params := fleet.NewCreateWorkstationParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithStoreID(storeID)
	params.WithBody(&models.CreateFleetWorkstation{Name: &name})

	res, err := client.Fleet.CreateWorkstation(params)

	if err != nil {
		log.Printf("Failed to create workstation %q", err)

		return err
	}

	components := strings.Split(res.Location, "/")
	workstationId := components[len(components)-1]
	d.SetId(workstationId)

	return nil
}

func resourceWorkstationRead(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Get("country_id").(string)
	storeID := d.Get("store_id").(string)

	params := fleet.NewGetWorkstationByIDParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithStoreID(storeID)
	params.WithWorkstationID(d.Id())
	res, err := client.Fleet.GetWorkstationByID(params)

	if err != nil {
		log.Printf("Failed to find workstation %q", err)
		d.SetId("")

		return nil
	}

	d.Set("name", res.Payload.Name)

	return nil
}

func resourceWorkstationUpdate(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Get("country_id").(string)
	storeID := d.Get("store_id").(string)
	name := d.Get("name").(string)

	params := fleet.NewUpdateWorkstationParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithStoreID(storeID)
	params.WithBody(&models.UpdateFleetWorkstation{Name: &name})

	_, err := client.Fleet.UpdateWorkstation(params)

	if err != nil {
		log.Printf("Failed to find workstation %q", err)
		d.SetId("")

		return nil
	}

	return nil
}

func resourceWorkstationDelete(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Get("country_id").(string)
	storeID := d.Get("store_id").(string)

	params := fleet.NewDeleteWorkstationParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithStoreID(storeID)
	params.WithWorkstationID(d.Id())

	_, err := client.Fleet.DeleteWorkstation(params)

	if err != nil {
		log.Printf("Failed to find workstation %q", err)

		return nil
	}

	return nil
}
