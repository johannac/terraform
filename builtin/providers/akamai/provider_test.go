package akamai

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"akamai": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	// if v := os.Getenv("ALICLOUD_ACCESS_KEY"); v == "" {
	//  t.Fatal("ALICLOUD_ACCESS_KEY must be set for acceptance tests")
	// }
	// if v := os.Getenv("ALICLOUD_SECRET_KEY"); v == "" {
	//  t.Fatal("ALICLOUD_SECRET_KEY must be set for acceptance tests")
	// }
	// if v := os.Getenv("ALICLOUD_REGION"); v == "" {
	//  log.Println("[INFO] Test: Using cn-beijing as test region")
	//  os.Setenv("ALICLOUD_REGION", "cn-beijing")
	// }
}
