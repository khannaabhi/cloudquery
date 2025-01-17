// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"

func Armdatadog() []*Table {
	tables := []*Table{
		{
			NewFunc:        armdatadog.NewMarketplaceAgreementsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/agreements",
			Namespace:      "Microsoft.Datadog",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Datadog)`,
			Pager:          `NewListPager`,
			ResponseStruct: "MarketplaceAgreementsClientListResponse",
		},
		{
			NewFunc:        armdatadog.NewMonitorsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/monitors",
			Namespace:      "Microsoft.Datadog",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Datadog)`,
			Pager:          `NewListPager`,
			ResponseStruct: "MonitorsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdatadog())
}
