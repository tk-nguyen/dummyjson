// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	dummyjson "demo.null/dummy"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &DummyProvider{}

//var _ provider.ProviderWithFunctions = &ScaffoldingProvider{}

// ScaffoldingProvider defines the provider implementation.
type DummyProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ScaffoldingProviderModel describes the provider data model.
type DummyProviderModel struct {
	Url types.String `tfsdk:"url"`
}

func (p *DummyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dummy"
	resp.Version = p.version
}

func (p *DummyProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				MarkdownDescription: "URL of the DummyJSON",
				Required:            true,
			},
		},
	}
}

func (p *DummyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data DummyProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }
	if data.Url.IsNull() {
		resp.Diagnostics.AddError("Empty URL provided!", "Please provide a valid URL to DummyJSON")
		return
	}
	if data.Url.ValueString() == "" {
		resp.Diagnostics.AddError("Empty URL provided!", "Please provide a valid URL to DummyJSON")
		return
	}

	// Example client configuration for data sources and resources
	client := dummyjson.NewDummyClient(data.Url.ValueString())
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *DummyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProductResource,
	}
}

func (p *DummyProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProductsDataSource,
		NewProductDataSource,
	}
}

func (p *DummyProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewExampleFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DummyProvider{
			version: version,
		}
	}
}
