// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAllProductsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		// PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + testAccAllProductsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.dummy_products.test", "products.0.id", "1"),
					resource.TestCheckResourceAttr("data.dummy_products.test", "products.#", "194"),
				),
			},
		},
	})
}

const testAccAllProductsDataSourceConfig = `
data "dummy_products" "test" {}
`
