package cloudflare

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudflareWAFRule_Import(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	ruleID := "100000"
	name := generateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudflareWAFRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareWAFRuleConfig(zoneID, ruleID, "block", name),
			},
			{
				ResourceName:        "cloudflare_waf_rule." + name,
				ImportStateIdPrefix: fmt.Sprintf("%s/", zoneID),
				ImportState:         true,
				ImportStateVerify:   true,
			},
		},
	})
}
