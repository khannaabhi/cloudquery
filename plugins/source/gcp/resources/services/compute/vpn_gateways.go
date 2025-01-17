// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/compute/apiv1"
)

func VpnGateways() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_vpn_gateways",
		Resolver:  fetchVpnGateways,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "label_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LabelFingerprint"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "stack_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackType"),
			},
			{
				Name:     "vpn_interfaces",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpnInterfaces"),
			},
		},
	}
}

func fetchVpnGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListVpnGatewaysRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewVpnGatewaysRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.AggregatedList(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp.Value.VpnGateways

	}
	return nil
}
