// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	dummyjson "demo.null/dummy"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ProductResource{}
var _ resource.ResourceWithImportState = &ProductResource{}

func NewProductResource() resource.Resource {
	return &ProductResource{}
}

// ProductResource defines the resource implementation.
type ProductResource struct {
	client *dummyjson.DummyClient
}

// ProductResourceModel describes the resource data model.
type ProductResourceModel struct {
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

func (r *ProductResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product"
}

func (r *ProductResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Create a new product on DummyJSON",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The ID of the product",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"title": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The title of the product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The description of the product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"category": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "The category of the product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"price": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The price of the product",
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"discount_percentage": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The discount percentage of the product",
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"rating": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The user rating of the product",
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"stock": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The remaining stock of the product",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "List of tags of the product",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"brand": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The brand of the product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sku": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The SKU of the product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"weight": schema.Float64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The weight of the product",
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"dimensions": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"width": schema.Float64Attribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Width of the product",
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"height": schema.Float64Attribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Height of the product",
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"depth": schema.Float64Attribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Depth of the product",
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
				},
				Computed:            true,
				MarkdownDescription: "The dimension of the product",
			},
			"warranty_info": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The warranty information of the product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"shipping_info": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The shipping information of the product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"availability_status": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The availability status of the product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"reviews": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rating": schema.Int64Attribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Product rating of this user",
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"comment": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Product comment of this user",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"date": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Review date of this user",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"reviewer_name": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Name of this reviewer",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"reviewer_email": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							MarkdownDescription: "Email of this reviewer",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"minimum_order_quantity": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The minimum order quantity of the product",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"thumbnail": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Thumbnail of this product",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"images": schema.ListAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Product images",
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ProductResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*dummyjson.DummyClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dummyjson.DummyClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ProductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ProductResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	configured := dummyjson.Product{
		Id:                   int(data.Id.ValueInt64()),
		Title:                data.Title.ValueString(),
		Description:          data.Description.ValueString(),
		Category:             data.Category.ValueString(),
		Price:                data.Price.ValueFloat64(),
		DiscountPercentage:   data.DiscountPercentage.ValueFloat64(),
		Rating:               data.Rating.ValueFloat64(),
		Stock:                uint(data.Stock.ValueInt64()),
		Brand:                data.Brand.ValueString(),
		Sku:                  data.Sku.ValueString(),
		Weight:               data.Weight.ValueFloat64(),
		WarrantyInfo:         data.WarrantyInfo.ValueString(),
		ShippingInfo:         data.ShippingInfo.ValueString(),
		AvailabilityStatus:   data.AvailabilityStatus.ValueString(),
		ReturnPolicy:         data.ReturnPolicy.ValueString(),
		MinimumOrderQuantity: uint(data.MinimumOrderQuantity.ValueInt64()),
		Thumbnail:            data.Thumbnail.ValueString(),
	}

	resp.Diagnostics.Append(data.Tags.ElementsAs(ctx, &configured.Tags, true)...)
	resp.Diagnostics.Append(data.Dimensions.As(ctx, &configured.Dimensions, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})...)
	resp.Diagnostics.Append(data.Reviews.ElementsAs(ctx, &configured.Reviews, true)...)
	resp.Diagnostics.Append(data.Images.ElementsAs(ctx, &configured.Images, true)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	product, err := r.client.UploadProduct(configured)
	if err != nil {
		resp.Diagnostics.AddError("Error creating product", fmt.Sprintf("Unable to create a new product, got error: %s", err))
		return
	}
	if resp.Diagnostics.HasError() {
		return
	}
	var diags diag.Diagnostics
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
	tflog.Trace(ctx, "Create a new product at DummyJSON")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ProductResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	product, err := r.client.GetProduct(int(data.Id.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError("Error getting product", fmt.Sprintf("Unable to read product from DummyJSON, got error: %s", err))
		return
	}
	var diags diag.Diagnostics
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

	if resp.Diagnostics.HasError() {
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ProductResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ProductResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (r *ProductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
