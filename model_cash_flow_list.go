/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Cash flows
type CashFlowList struct {
	// Array of cash flows
	CashFlows []CashFlow `json:"cashFlows"`
	// The total income
	TotalIncome float32 `json:"totalIncome"`
	// The total spending
	TotalSpending float32 `json:"totalSpending"`
	// The total balance
	TotalBalance float32 `json:"totalBalance"`
}
