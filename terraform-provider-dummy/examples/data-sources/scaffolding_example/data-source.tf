# Get all products from DummyJSON
data "dummy_products" "all_products" {}

# Get a single product from DummyJSON
data "dummy_product" "single_product" {
  # Only id is required
  # valid values: 1 - 194
  id = 123
}
