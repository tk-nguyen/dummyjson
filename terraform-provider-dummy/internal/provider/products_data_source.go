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
var _ datasource.DataSource = &ProductsDataSource{}

func NewProductsDataSource() datasource.DataSource {
	return &ProductsDataSource{}
}

// ExampleDataSource defines the data source implementation.
type ProductsDataSource struct {
	client *dummyjson.DummyClient
}
type ProductsDataSourceModel struct {
	Products types.List `tfsdk:"products"`
}

func (d *ProductsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_products"
}

func (d *ProductsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Products data source",

		Attributes: map[string]schema.Attribute{
			"products": schema.ListNestedAttribute{
				Computed:            true,
				MarkdownDescription: "List of all products in DummyJSON",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "The ID of the product",
						},
						"title": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The title of the product",
						},
						"description": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The description of the product",
						},
						"category": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The category of the product",
						},
						"price": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "The price of the product",
						},
						"discount_percentage": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "The discount percentage of the product",
						},
						"rating": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "The user rating of the product",
						},
						"stock": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "The remaining stock of the product",
						},
						"tags": schema.ListAttribute{
							Computed:            true,
							ElementType:         types.StringType,
							MarkdownDescription: "List of tags of the product",
						},
						"brand": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The brand of the product",
						},
						"sku": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The SKU of the product",
						},
						"weight": schema.Float64Attribute{
							Computed:            true,
							MarkdownDescription: "The weight of the product",
						},
						"dimensions": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"width": schema.Float64Attribute{
									Computed:            true,
									MarkdownDescription: "Width of the product",
								},
								"height": schema.Float64Attribute{
									Computed:            true,
									MarkdownDescription: "Height of the product",
								},
								"depth": schema.Float64Attribute{
									Computed:            true,
									MarkdownDescription: "Depth of the product",
								},
							},
							Computed: true,

							MarkdownDescription: "The dimension of the product",
						},
						"warranty_info": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The warranty information of the product",
						},
						"shipping_info": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The shipping information of the product",
						},
						"availability_status": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The availability status of the product",
						},
						"reviews": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"rating": schema.Int64Attribute{
										Computed:            true,
										MarkdownDescription: "Product rating of this user",
									},
									"comment": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Product comment of this user",
									},
									"date": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Review date of this user",
									},
									"reviewer_name": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Name of this reviewer",
									},
									"reviewer_email": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: "Email of this reviewer",
									},
								},
							},
							Computed:            true,
							MarkdownDescription: "The reviews of the product",
						},
						"return_policy": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The return policy of the product",
						},
						"minimum_order_quantity": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "The minimum order quantity of the product",
						},
						"thumbnail": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Thumbnail of this product",
						},
						"images": schema.ListAttribute{
							ElementType:         types.StringType,
							Computed:            true,
							MarkdownDescription: "Product images",
						},
					},
				},
			},
		},
	}
}

func (d *ProductsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*dummyjson.DummyClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ProductsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ProductsDataSourceModel
	// This is for converting Go types to Terraform types
	var diags diag.Diagnostics
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	products, err := d.client.GetProducts()
	if err != nil {
		resp.Diagnostics.AddError("DummyClient Error", fmt.Sprintf("Unable to get products from DummyJSON , got error: %s", err))
		return
	}
	// Product data with terraform types
	productDataTf := make([]types.Object, 0)
	for _, p := range products {
		prod := ProductModel{
			Id:                   types.Int64Value(int64(p.Id)),
			Title:                types.StringValue(p.Title),
			Description:          types.StringValue(p.Description),
			Category:             types.StringValue(p.Category),
			Price:                types.Float64Value(p.Price),
			DiscountPercentage:   types.Float64Value(p.Price),
			Rating:               types.Float64Value(p.Rating),
			Stock:                types.Int64Value(int64(p.Stock)),
			Brand:                types.StringValue(p.Brand),
			Sku:                  types.StringValue(p.Sku),
			Weight:               types.Float64Value(p.Weight),
			WarrantyInfo:         types.StringValue(p.WarrantyInfo),
			ShippingInfo:         types.StringValue(p.ShippingInfo),
			AvailabilityStatus:   types.StringValue(p.AvailabilityStatus),
			ReturnPolicy:         types.StringValue(p.ReturnPolicy),
			MinimumOrderQuantity: types.Int64Value(int64(p.MinimumOrderQuantity)),
			Thumbnail:            types.StringValue(p.ShippingInfo),
		}
		prod.Tags, diags = types.ListValueFrom(ctx, types.StringType, p.Tags)
		resp.Diagnostics.Append(diags...)
		prod.Images, diags = types.ListValueFrom(ctx, types.StringType, p.Images)
		resp.Diagnostics.Append(diags...)
		reviews := make([]types.Object, 0)
		for _, review := range p.Reviews {
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
		// Create List[Object{ReviewModelType}] from go type
		prod.Reviews, diags = types.ListValueFrom(ctx, types.ObjectType{
			AttrTypes: ReviewModelType,
		}, reviews)
		resp.Diagnostics.Append(diags...)
		dimension := DimensionModel{
			Width:  types.Float64Value(p.Dimensions.Width),
			Height: types.Float64Value(p.Dimensions.Height),
			Depth:  types.Float64Value(p.Dimensions.Depth),
		}
		// Create DimensionModelType from go type
		prod.Dimensions, diags = types.ObjectValueFrom(ctx, DimensionModelType, dimension)
		resp.Diagnostics.Append(diags...)
		// Create ProductModelType from go type
		prodObj, diags := types.ObjectValueFrom(ctx, ProductModelType, prod)
		resp.Diagnostics.Append(diags...)
		productDataTf = append(productDataTf, prodObj)
	}

	data.Products, diags = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ProductModelType}, productDataTf)
	resp.Diagnostics.Append(diags...)
	// If there's error when converting between Go and Terraform types
	// return early
	if resp.Diagnostics.HasError() {
		return
	}

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "Successfully get all products from DummyJSON!")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
