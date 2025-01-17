// Code generated by codegen2; DO NOT EDIT.
package devops

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PipelineTemplateDefinitions() *schema.Table {
	return &schema.Table{
		Name:      "azure_devops_pipeline_template_definitions",
		Resolver:  fetchPipelineTemplateDefinitions,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DevOps),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "inputs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Inputs"),
			},
		},
	}
}

func fetchPipelineTemplateDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armdevops.NewPipelineTemplateDefinitionsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
