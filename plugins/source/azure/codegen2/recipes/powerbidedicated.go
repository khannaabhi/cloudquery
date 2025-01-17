// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"

func init() {
	tables := []Table{
		{
			Service:        "armpowerbidedicated",
			Name:           "capacities",
			Struct:         &armpowerbidedicated.DedicatedCapacity{},
			ResponseStruct: &armpowerbidedicated.CapacitiesClientListResponse{},
			Client:         &armpowerbidedicated.CapacitiesClient{},
			ListFunc:       (&armpowerbidedicated.CapacitiesClient{}).NewListPager,
			NewFunc:        armpowerbidedicated.NewCapacitiesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/capacities",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_PowerBIDedicated)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
