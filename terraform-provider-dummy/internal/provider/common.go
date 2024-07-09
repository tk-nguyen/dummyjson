// This file contain common models for DummyJSON

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductModel struct {
	Id                   types.Int64   `tfsdk:"id"`
	Title                types.String  `tfsdk:"title"`
	Description          types.String  `tfsdk:"description"`
	Category             types.String  `tfsdk:"category"`
	Price                types.Float64 `tfsdk:"price"`
	DiscountPercentage   types.Float64 `tfsdk:"discount_percentage"`
	Rating               types.Float64 `tfsdk:"rating"`
	Stock                types.Int64   `tfsdk:"stock"`
	Tags                 types.List    `tfsdk:"tags"`
	Brand                types.String  `tfsdk:"brand"`
	Sku                  types.String  `tfsdk:"sku"`
	Weight               types.Float64 `tfsdk:"weight"`
	Dimensions           types.Object  `tfsdk:"dimensions"`
	WarrantyInfo         types.String  `tfsdk:"warranty_info"`
	ShippingInfo         types.String  `tfsdk:"shipping_info"`
	AvailabilityStatus   types.String  `tfsdk:"availability_status"`
	Reviews              types.List    `tfsdk:"reviews"`
	ReturnPolicy         types.String  `tfsdk:"return_policy"`
	MinimumOrderQuantity types.Int64   `tfsdk:"minimum_order_quantity"`
	Thumbnail            types.String  `tfsdk:"thumbnail"`
	Images               types.List    `tfsdk:"images"`
}

var ProductModelType = map[string]attr.Type{
	"id":                  types.Int64Type,
	"title":               types.StringType,
	"description":         types.StringType,
	"category":            types.StringType,
	"price":               types.Float64Type,
	"discount_percentage": types.Float64Type,
	"rating":              types.Float64Type,
	"stock":               types.Int64Type,
	"tags": types.ListType{
		ElemType: types.StringType,
	},
	"brand":                  types.StringType,
	"sku":                    types.StringType,
	"weight":                 types.Float64Type,
	"dimensions":             types.ObjectType{AttrTypes: DimensionModelType},
	"warranty_info":          types.StringType,
	"shipping_info":          types.StringType,
	"availability_status":    types.StringType,
	"reviews":                types.ListType{ElemType: types.ObjectType{AttrTypes: ReviewModelType}},
	"return_policy":          types.StringType,
	"minimum_order_quantity": types.Int64Type,
	"thumbnail":              types.StringType,
	"images":                 types.ListType{ElemType: types.StringType},
}

type ReviewModel struct {
	Rating        types.Int64  `tfsdk:"rating"`
	Comment       types.String `tfsdk:"comment"`
	Date          types.String `tfsdk:"date"`
	ReviewerName  types.String `tfsdk:"reviewer_name"`
	ReviewerEmail types.String `tfsdk:"reviewer_email"`
}

var ReviewModelType = map[string]attr.Type{
	"rating":         types.Int64Type,
	"comment":        types.StringType,
	"date":           types.StringType,
	"reviewer_name":  types.StringType,
	"reviewer_email": types.StringType,
}

type DimensionModel struct {
	Width  types.Float64 `tfsdk:"width"`
	Height types.Float64 `tfsdk:"height"`
	Depth  types.Float64 `tfsdk:"depth"`
}

var DimensionModelType = map[string]attr.Type{
	"width":  types.Float64Type,
	"height": types.Float64Type,
	"depth":  types.Float64Type,
}
