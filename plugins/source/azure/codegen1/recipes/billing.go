// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"

func Armbilling() []*Table {
	tables := []*Table{
		{
			NewFunc:        armbilling.NewAccountsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling",
			URL:            "/providers/Microsoft.Billing/billingAccounts",
			Namespace:      "Microsoft.Billing",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Billing)`,
			Pager:          `NewListPager`,
			ResponseStruct: "AccountsClientListResponse",
		},
		{
			NewFunc:        armbilling.NewEnrollmentAccountsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling",
			URL:            "/providers/Microsoft.Billing/enrollmentAccounts",
			Namespace:      "Microsoft.Billing",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Billing)`,
			Pager:          `NewListPager`,
			ResponseStruct: "EnrollmentAccountsClientListResponse",
		},
		{
			NewFunc:        armbilling.NewPeriodsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods",
			Namespace:      "Microsoft.Billing",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Billing)`,
			Pager:          `NewListPager`,
			ResponseStruct: "PeriodsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armbilling())
}
