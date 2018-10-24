package fleetmanager

import (
	"errors"
	"testing"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFleetManagerStore(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testStoreDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFleetManagerStore,
				Check: resource.ComposeTestCheckFunc(testStoreExists("fleetmanager_store.test"),
					resource.TestCheckResourceAttr("fleetmanager_store.test", "name", "A Test Store"),
					resource.TestCheckResourceAttr("fleetmanager_store.test", "city", "City"),
					resource.TestCheckResourceAttr("fleetmanager_store.test", "postcode", "POST CODE"),
					resource.TestCheckResourceAttr("fleetmanager_store.test", "street", "Test St"),
					resource.TestCheckResourceAttr("fleetmanager_store.test", "contact_person.first_name", "John"),
					resource.TestCheckResourceAttr("fleetmanager_store.test", "contact_person.last_name", "Doe"),
					resource.TestCheckResourceAttr("fleetmanager_store.test", "contact_person.email", "john.doe@example.com"),
					resource.TestCheckResourceAttr("fleetmanager_store.test", "contact_person.phone", "+44 020 7946 0123"),
				),
			},
		},
	})
}

func testStoreDestroy(s *terraform.State) error {
	tenant := testAccProvider.Meta().(*Tenant)
	store := s.RootModule().Resources["fleetmanager_store.test"].Primary
	_, err := fetchStore(tenant, store)

	if err == nil {
		return errors.New("store is still present")
	}

	return nil
}

func testStoreExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		tenant := testAccProvider.Meta().(*Tenant)
		store := s.RootModule().Resources["fleetmanager_store.test"].Primary
		_, err := fetchStore(tenant, store)

		if err != nil {
			return err
		}

		return nil
	}
}

func fetchStore(tenant *Tenant, store *terraform.InstanceState) (*fleet.GetStoreByIDOK, error) {
	client := tenant.Client
	brandID, ok := store.Attributes["brand_id"]
	if !ok {
		return nil, errors.New("brandID not found on store resource")
	}

	countryID, ok := store.Attributes["country_id"]
	if !ok {
		return nil, errors.New("countryID not found on store resource")
	}

	params := fleet.NewGetStoreByIDParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithStoreID(store.ID)

	return client.Fleet.GetStoreByID(params)
}

const testAccCheckFleetManagerStore = `
resource "fleetmanager_brand" "test" {
	name = "A Test Brand"
}

resource "fleetmanager_country" "test" {
	brand_id = "${fleetmanager_brand.test.id}"

	name = "Unique Country"
	iso_code = "abcdef"
}

resource "fleetmanager_store" "test" {
	brand_id = "${fleetmanager_brand.test.id}"
	country_id = "${fleetmanager_country.test.id}"

	name = "A Test Store"
	city = "City"
	postcode = "POST CODE"
	street = "Test St"

	contact_person {
		first_name = "John"
		last_name = "Doe"
		email = "john.doe@example.com"
		phone = "+44 020 7946 0123"
	}
}
`
