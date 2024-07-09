// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSingleProductDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		// PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + testAccSingleProductDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.dummy_product.test", "id", "123"),
					resource.TestCheckResourceAttr("data.dummy_product.test", "title", "iPhone 13 Pro"),
					resource.TestCheckResourceAttr("data.dummy_product.test", "sku", "YGQKHPGK"),
					resource.TestCheckResourceAttr("data.dummy_product.test", "reviews.0.reviewer_name", "Aria Roberts"),
					resource.TestCheckResourceAttr("data.dummy_product.test", "reviews.0.reviewer_email", "aria.roberts@x.dummyjson.com"),
				),
			},
		},
	})
}

const testAccSingleProductDataSourceConfig = `
data "dummy_product" "test" {
	id = 123 
}
`
