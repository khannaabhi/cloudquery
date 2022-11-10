// Code generated by codegen; DO NOT EDIT.

package ram

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resources",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_Resource.html`,
		Resolver:    fetchRamResources,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "resource_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceGroupArn"),
			},
			{
				Name:     "resource_region_scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceRegionScope"),
			},
			{
				Name:     "resource_share_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareArn"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusMessage"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}