package fleetmanager

import (
	"errors"
	"testing"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFleetManagerWorkstation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testWorkstationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFleetManagerWorkstation,
				Check: resource.ComposeTestCheckFunc(testWorkstationExists("fleetmanager_workstation.test"),
					resource.TestCheckResourceAttr("fleetmanager_workstation.test", "name", "A Test Workstation"),
				),
			},
		},
	})
}

func testWorkstationDestroy(s *terraform.State) error {
	tenant := testAccProvider.Meta().(*Tenant)
	workstation := s.RootModule().Resources["fleetmanager_workstation.test"].Primary
	_, err := fetchWorkstation(tenant, workstation)

	if err == nil {
		return errors.New("workstation is still present")
	}

	return nil
}

func testWorkstationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		tenant := testAccProvider.Meta().(*Tenant)
		workstation := s.RootModule().Resources["fleetmanager_workstation.test"].Primary
		_, err := fetchWorkstation(tenant, workstation)

		if err != nil {
			return err
		}

		return nil
	}
}

func fetchWorkstation(tenant *Tenant, workstation *terraform.InstanceState) (*fleet.GetWorkstationByIDOK, error) {
	client := tenant.Client
	brandID, ok := workstation.Attributes["brand_id"]
	if !ok {
		return nil, errors.New("brandID not found on workstation resource")
	}

	countryID, ok := workstation.Attributes["country_id"]
	if !ok {
		return nil, errors.New("countryID not found on workstation resource")
	}

	storeID, ok := workstation.Attributes["store_id"]
	if !ok {
		return nil, errors.New("storeID not found on workstation resource")
	}

	params := fleet.NewGetWorkstationByIDParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(countryID)
	params.WithStoreID(storeID)
	params.WithWorkstationID(workstation.ID)

	return client.Fleet.GetWorkstationByID(params)
}

const testAccCheckFleetManagerWorkstation = `
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

resource "fleetmanager_workstation" "test" {
	brand_id = "${fleetmanager_brand.test.id}"
	country_id = "${fleetmanager_country.test.id}"
	store_id = "${fleetmanager_store.test.id}"

	name = "A Test Workstation"
}
`
