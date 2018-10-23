package fleetmanager

import (
	"errors"
	"testing"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFleetManagerCountry(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCountryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFleetManagerCountry,
				Check: resource.ComposeTestCheckFunc(
					testCountryExists("fleetmanager_country.test"),
					resource.TestCheckResourceAttr("fleetmanager_country.test",
						"name", "Unique Country"),
					resource.TestCheckResourceAttr("fleetmanager_country.test",
						"iso_code", "abcdef"),
				),
			},
		},
	})
}

func testCountryDestroy(s *terraform.State) error {
	tenant := testAccProvider.Meta().(*Tenant)
	if err := testCountryDestroyHelper(s, tenant); err != nil {
		return err
	}

	return nil
}

func testCountryDestroyHelper(s *terraform.State, tenant *Tenant) error {
	client := tenant.Client
	country := s.RootModule().Resources["fleetmanager_country.test"]

	brandID, ok := country.Primary.Attributes["brand_id"]
	if !ok {
		return errors.New("brandID not found on country resource")
	}

	params := fleet.NewGetCountryByIDParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(country.Primary.ID)
	_, err := client.Fleet.GetCountryByID(params)

	if _, ok := err.(*fleet.GetCountryByIDNotFound); !ok {
		return err
	}

	return nil
}

func testCountryExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		tenant := testAccProvider.Meta().(*Tenant)
		if err := testCountryExistsHelper(s, tenant); err != nil {
			return err
		}

		return nil
	}
}

func testCountryExistsHelper(s *terraform.State, tenant *Tenant) error {
	client := tenant.Client
	country := s.RootModule().Resources["fleetmanager_country.test"]

	brandID, ok := country.Primary.Attributes["brand_id"]
	if !ok {
		return errors.New("brandID not found on country resource")
	}

	params := fleet.NewGetCountryByIDParams()
	params.WithTenantID(tenant.ID)
	params.WithBrandID(brandID)
	params.WithCountryID(country.Primary.ID)

	_, err := client.Fleet.GetCountryByID(params)
	if err != nil {
		return err
	}

	return nil
}

const testAccCheckFleetManagerCountry = `
resource "fleetmanager_brand" "test" {
	name = "A Test Brand"
}

resource "fleetmanager_country" "test" {
	brand_id = "${fleetmanager_brand.test.id}"

	name = "Unique Country"
	iso_code = "abcdef"
}
`
