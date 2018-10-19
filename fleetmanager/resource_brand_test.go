package fleetmanager

import (
	"testing"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_brand"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"fleetmanager": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {} // should use to test env vars are set

func TestAccFleetManagerBrand(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testBrandDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFleetManagerBrand,
				Check: resource.ComposeTestCheckFunc(
					testBrandExists("fleetmanager_brand.test"),
					resource.TestCheckResourceAttr("fleetmanager_brand.test", "name", "A Test Brand"),
				),
			},
		},
	})
}

func testBrandDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Fleetmanager)
	if err := testBrandDestroyHelper(s, client); err != nil {
		return err
	}
	return nil
}

func testBrandDestroyHelper(s *terraform.State, client *client.Fleetmanager) error {
	for _, r := range s.RootModule().Resources {
		params := fleet_brand.NewGetFleetTenantTenantIDBrandBrandIDParams()
		params.WithBrandID(r.Primary.ID)
		params.WithTenantID("hm")
		_, err := client.FleetBrand.GetFleetTenantTenantIDBrandBrandID(params)

		if _, ok := err.(*fleet_brand.GetFleetTenantTenantIDBrandBrandIDNotFound); !ok {
			return err
		}
	}

	return nil
}

func testBrandExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*client.Fleetmanager)
		if err := testBrandExistsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

func testBrandExistsHelper(s *terraform.State, client *client.Fleetmanager) error {
	for _, r := range s.RootModule().Resources {
		params := fleet_brand.NewGetFleetTenantTenantIDBrandBrandIDParams()
		params.WithBrandID(r.Primary.ID)
		params.WithTenantID("hm")
		_, err := client.FleetBrand.GetFleetTenantTenantIDBrandBrandID(params)

		if err != nil {
			return err
		}
	}

	return nil
}

const testAccCheckFleetManagerBrand = `
resource "fleetmanager_brand" "test" {
	name = "A Test Brand"
}
`
