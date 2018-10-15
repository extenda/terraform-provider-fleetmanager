package fleetmanager

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("FM_HOST"); v == "" {
		t.Fatal("FM_HOST must be set for acceptance tests")
	}

	if v := os.Getenv("FM_TENANT"); v == "" {
		t.Fatal("FM_TENANT must be set for acceptance tests")
	}
}
