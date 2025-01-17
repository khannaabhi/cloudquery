// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"

func init() {
	tables := []Table{
		{
			Service:        "armbatch",
			Name:           "account",
			Struct:         &armbatch.Account{},
			ResponseStruct: &armbatch.AccountClientListResponse{},
			Client:         &armbatch.AccountClient{},
			ListFunc:       (&armbatch.AccountClient{}).NewListPager,
			NewFunc:        armbatch.NewAccountClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Batch/batchAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Batch)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
