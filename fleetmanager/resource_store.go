package fleetmanager

import (
	"log"
	"strings"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceStoreCreate,
		Read:   resourceStoreRead,
		Update: resourceStoreUpdate,
		Delete: resourceStoreDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"city": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"postcode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"street": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"country_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"brand_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"contact_person": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"last_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"phone": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceStoreCreate(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Get("country_id").(string)

	contactPersonMap := d.Get("contact_person").(map[string]interface{})
	email := contactPersonMap["email"].(string)
	firstName := contactPersonMap["first_name"].(string)
	lastName := contactPersonMap["last_name"].(string)
	phone := contactPersonMap["phone"].(string)

	contactPerson := &models.FleetStoreContactPerson{
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
		Phone:     &phone,
	}

	name := d.Get("name").(string)
	city := d.Get("city").(string)
	postCode := d.Get("postcode").(string)
	street := d.Get("street").(string)

	store := &models.CreateFleetStore{
		Name:          &name,
		City:          &city,
		PostCode:      &postCode,
		Street:        &street,
		ContactPerson: contactPerson,
	}

	params := fleet.NewCreateStoreParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithBody(store)

	res, err := client.Fleet.CreateStore(params)

	if err != nil {
		log.Printf("Failed to create store %q", err)

		return err
	}

	components := strings.Split(res.Location, "/")
	storeId := components[len(components)-1]
	d.SetId(storeId)

	return nil
}

func resourceStoreRead(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Get("country_id").(string)

	params := fleet.NewGetStoreByIDParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithStoreID(d.Id())
	res, err := client.Fleet.GetStoreByID(params)

	if err != nil {
		log.Printf("Failed to find store %q", err)
		d.SetId("")

		return nil
	}

	d.Set("name", res.Payload.Name)
	d.Set("city", res.Payload.City)
	d.Set("postcode", res.Payload.PostCode)
	d.Set("street", res.Payload.Street)

	cp := res.Payload.ContactPerson
	d.Set("contact_person.first_name", cp.FirstName)
	d.Set("contact_person.last_name", cp.LastName)
	d.Set("contact_person.email", cp.Email)
	d.Set("contact_person.phone", cp.Phone)

	return nil
}

func resourceStoreUpdate(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Get("country_id").(string)

	contactPersonMap := d.Get("contact_person").(map[string]interface{})
	email := contactPersonMap["email"].(string)
	firstName := contactPersonMap["first_name"].(string)
	lastName := contactPersonMap["last_name"].(string)
	phone := contactPersonMap["phone"].(string)

	contactPerson := &models.FleetStoreContactPerson{
		Email:     &email,
		FirstName: &firstName,
		LastName:  &lastName,
		Phone:     &phone,
	}

	name := d.Get("name").(string)
	city := d.Get("city").(string)
	postCode := d.Get("postcode").(string)
	street := d.Get("street").(string)

	store := &models.UpdateFleetStore{
		Name:          &name,
		City:          &city,
		PostCode:      &postCode,
		Street:        &street,
		ContactPerson: contactPerson,
	}

	params := fleet.NewUpdateStoreParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithBody(store)

	_, err := client.Fleet.UpdateStore(params)

	if err != nil {
		log.Printf("Failed to find store %q", err)
		d.SetId("")

		return nil
	}

	return nil
}

func resourceStoreDelete(d *schema.ResourceData, m interface{}) error {
	tenant := m.(*Tenant)
	client := tenant.Client
	brandID := d.Get("brand_id").(string)
	countryID := d.Get("country_id").(string)

	params := fleet.NewDeleteStoreParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithStoreID(d.Id())

	_, err := client.Fleet.DeleteStore(params)

	if err != nil {
		log.Printf("Failed to find store %q", err)

		return nil
	}

	return nil
}
