// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"

func Armconfluent() []*Table {
	tables := []*Table{
		{
			NewFunc:        armconfluent.NewMarketplaceAgreementsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Confluent/agreements",
			Namespace:      "Microsoft.Confluent",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Confluent)`,
			Pager:          `NewListPager`,
			ResponseStruct: "MarketplaceAgreementsClientListResponse",
		},
		{
			NewFunc:        armconfluent.NewOrganizationOperationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent",
			URL:            "/providers/Microsoft.Confluent/operations",
			Namespace:      "Microsoft.Confluent",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Confluent)`,
			Pager:          `NewListPager`,
			ResponseStruct: "OrganizationOperationsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armconfluent())
}
