// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	dummyjson "demo.null/dummy"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ProductDataSource{}

type ProductDataSource struct {
	client *dummyjson.DummyClient
}

func NewProductDataSource() datasource.DataSource {
	return &ProductDataSource{}
}

type ProductDataSourceModel struct {
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

func (d *ProductDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product"
}

func (d *ProductDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Get information about a product",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The ID of the product",
			},
			"title": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The title of the product",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The description of the product",
			},
			"category": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "The category of the product",
			},
			"price": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The price of the product",
			},
			"discount_percentage": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The discount percentage of the product",
			},
			"rating": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The user rating of the product",
			},
			"stock": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The remaining stock of the product",
			},
			"tags": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "List of tags of the product",
			},
			"brand": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The brand of the product",
			},
			"sku": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The SKU of the product",
			},
			"weight": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The weight of the product",
			},
			"dimensions": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"width": schema.Float64Attribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Width of the product",
					},
					"height": schema.Float64Attribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Height of the product",
					},
					"depth": schema.Float64Attribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Depth of the product",
					},
				},
				Computed: true,

				MarkdownDescription: "The dimension of the product",
			},
			"warranty_info": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The warranty information of the product",
			},
			"shipping_info": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The shipping information of the product",
			},
			"availability_status": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The availability status of the product",
			},
			"reviews": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rating": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Product rating of this user",
						},
						"comment": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Product comment of this user",
						},
						"date": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Review date of this user",
						},
						"reviewer_name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Name of this reviewer",
						},
						"reviewer_email": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Email of this reviewer",
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "The reviews of the product",
			},
			"return_policy": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The return policy of the product",
			},
			"minimum_order_quantity": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The minimum order quantity of the product",
			},
			"thumbnail": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Thumbnail of this product",
			},
			"images": schema.ListAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Product images",
			},
		},
	}
}

func (d *ProductDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*dummyjson.DummyClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *dummyjson.DummyClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ProductDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ProductDataSourceModel

	// This is for converting Go types to Terraform types
	var diags diag.Diagnostics

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	product, err := d.client.GetProduct(int(data.Id.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError("DummyClient Error", fmt.Sprintf("Unable to get products from DummyJSON , got error: %s", err))
		return
	}
	// Update current data with the ones from DummyJSON
	data.Id = types.Int64Value(int64(product.Id))
	data.Title = types.StringValue(product.Title)
	data.Description = types.StringValue(product.Description)
	data.Category = types.StringValue(product.Category)
	data.Id = types.Int64Value(int64(product.Id))
	data.Price = types.Float64Value(product.Price)
	data.DiscountPercentage = types.Float64Value(product.DiscountPercentage)
	data.Rating = types.Float64Value(product.Rating)
	data.Stock = types.Int64Value(int64(product.Stock))
	data.Tags, diags = types.ListValueFrom(ctx, types.StringType, product.Tags)
	resp.Diagnostics.Append(diags...)
	data.Brand = types.StringValue(product.Brand)
	data.Sku = types.StringValue(product.Sku)
	data.Weight = types.Float64Value(product.Weight)

	dimension := DimensionModel{
		Width:  types.Float64Value(product.Dimensions.Width),
		Height: types.Float64Value(product.Dimensions.Height),
		Depth:  types.Float64Value(product.Dimensions.Depth),
	}
	data.Dimensions, diags = types.ObjectValueFrom(ctx, DimensionModelType, dimension)
	resp.Diagnostics.Append(diags...)
	data.WarrantyInfo = types.StringValue(product.WarrantyInfo)
	data.ShippingInfo = types.StringValue(product.ShippingInfo)
	data.AvailabilityStatus = types.StringValue(product.AvailabilityStatus)
	data.ReturnPolicy = types.StringValue(product.ReturnPolicy)
	data.MinimumOrderQuantity = types.Int64Value(int64(product.MinimumOrderQuantity))
	data.Thumbnail = types.StringValue(product.Thumbnail)
	data.Images, diags = types.ListValueFrom(ctx, types.StringType, product.Images)
	resp.Diagnostics.Append(diags...)
	reviews := make([]types.Object, 0)
	for _, review := range product.Reviews {
		rm := ReviewModel{
			Rating:        types.Int64Value(int64(review.Rating)),
			Comment:       types.StringValue(review.Comment),
			Date:          types.StringValue(review.Date.String()),
			ReviewerName:  types.StringValue(review.ReviewerName),
			ReviewerEmail: types.StringValue(review.ReviewerEmail),
		}
		reviewObj, diags := types.ObjectValueFrom(ctx, ReviewModelType, rm)
		resp.Diagnostics.Append(diags...)
		reviews = append(reviews, reviewObj)
	}
	data.Reviews, diags = types.ListValueFrom(ctx, types.ObjectType{
		AttrTypes: ReviewModelType,
	}, reviews)
	resp.Diagnostics.Append(diags...)
	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "Successfully get product from DummyJSON!", map[string]interface{}{
		"Product ID": product.Id,
	})
	// If there's error when converting between Go and Terraform types
	// return early
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
