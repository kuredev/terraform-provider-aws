package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
)

type methodDataSource struct {
	framework.DataSourceWithModel[methodDataSourceModel]
}

func (d *methodDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"http_method": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

func (d *methodDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data methodDataSourceModel

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type methodDataSourceModel struct {
	httpMethod types.String `tfsdk:"http_method"`
}
