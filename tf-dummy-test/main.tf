terraform {
    required_providers {
      dummy = {
        source = "demo.null/terraform/dummy"
      }
    }
}

provider "dummy" {
    url = "https://dummyjson.com"
}

data "dummy_products" "prods" {}

data "dummy_product" "single_prod" {
  id = "100"

}
# Comment this out to prevent terminal exploding
# output "all_products" {
#     description = "All DummyJSON products"
#     value = data.dummy_products.prods.products
# }
output "single_product" {
    description = "Single DummyJSON product"
    value = data.dummy_product.single_prod
}