// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"

func init() {
	tables := []Table{
		{
			Service:        "armconfluent",
			Name:           "marketplace_agreements",
			Struct:         &armconfluent.AgreementResource{},
			ResponseStruct: &armconfluent.MarketplaceAgreementsClientListResponse{},
			Client:         &armconfluent.MarketplaceAgreementsClient{},
			ListFunc:       (&armconfluent.MarketplaceAgreementsClient{}).NewListPager,
			NewFunc:        armconfluent.NewMarketplaceAgreementsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Confluent/agreements",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Confluent)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}