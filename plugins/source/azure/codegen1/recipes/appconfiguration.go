// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"

func Armappconfiguration() []*Table {
	tables := []*Table{
		{
			NewFunc:        armappconfiguration.NewConfigurationStoresClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/configurationStores",
			Namespace:      "Microsoft.AppConfiguration",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_AppConfiguration)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ConfigurationStoresClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armappconfiguration())
}
