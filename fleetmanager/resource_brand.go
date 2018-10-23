package fleetmanager

import (
	"log"
	"strings"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet"
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
	tenant := m.(*Tenant)
	client := tenant.Client
	name := d.Get("name").(string)

	params := fleet.NewCreateBrandParams()
	params.WithTenantID(tenant.ID)
	params.WithBody(&models.CreateFleetBrand{Name: &name})

	res, err := client.Fleet.CreateBrand(params)

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
	tenant := m.(*Tenant)
	client := tenant.Client

	params := fleet.NewGetBrandByIDParams()
	params.WithBrandID(d.Id())
	params.WithTenantID(tenant.ID)
	res, err := client.Fleet.GetBrandByID(params)

	if err != nil {
		log.Printf("Failed to find brand %q", err)
		d.SetId("")

		return nil
	}

	d.Set("name", res.Payload.Name)

	return nil
}

func resourceBrandUpdate(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	name := d.Get("name").(string)

	params := fleet.NewUpdateBrandParams()
	params.WithBrandID(d.Id())
	params.WithTenantID(tenant.ID)
	params.WithBody(&models.UpdateFleetBrand{Name: &name})

	_, err := client.Fleet.UpdateBrand(params)

	if err != nil {
		log.Printf("Failed to find brand %q", err)
		d.SetId("")

		return nil
	}

	return nil
}

func resourceBrandDelete(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client

	params := fleet.NewDeleteBrandParams()
	params.WithBrandID(d.Id())
	params.WithTenantID(tenant.ID)

	_, err := client.Fleet.DeleteBrand(params)

	if err != nil {
		log.Printf("Failed to find brand %q", err)

		return nil
	}

	return nil
}
