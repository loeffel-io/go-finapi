/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Cash flow
type CashFlow struct {
	// Category of this cash flow. When null, then this is the cash flow of transactions that do not have a category.
	Category *Category `json:"category,omitempty"`
	// The total calculated income for the given category
	Income float32 `json:"income"`
	// The total calculated spending for the given category
	Spending float32 `json:"spending"`
	// The calculated balance for the given category
	Balance float32 `json:"balance"`
	// The total count of income transactions for the given category
	CountIncomeTransactions int32 `json:"countIncomeTransactions"`
	// The total count of spending transactions for the given category
	CountSpendingTransactions int32 `json:"countSpendingTransactions"`
	// The total count of all transactions for the given category
	CountAllTransactions int32 `json:"countAllTransactions"`
}