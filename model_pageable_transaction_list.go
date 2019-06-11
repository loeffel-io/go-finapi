/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for a page of transactions, with data about the total count of transactions and their balance
type PageableTransactionList struct {
	// Array of transactions (for the requested page)
	Transactions []Transaction `json:"transactions"`
	// Information for pagination
	Paging *Paging `json:"paging"`
	// The total income of all transactions (across all pages)
	Income float32 `json:"income"`
	// The total spending of all transactions (across all pages)
	Spending float32 `json:"spending"`
	// The total sum of all transactions (across all pages)
	Balance float32 `json:"balance"`
}