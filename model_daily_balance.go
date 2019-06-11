/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Balance data for a single day
type DailyBalance struct {
	// Date in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	Date string `json:"date"`
	// Calculated balance at the end of day (aggregation of all regarded accounts).
	Balance float32 `json:"balance"`
	// The sum of income of the given day (aggregation of all regarded accounts).
	Income float32 `json:"income"`
	// The sum of spending of the given day (aggregation of all regarded accounts).
	Spending float32 `json:"spending"`
	// Identifiers of the transactions for the given day
	Transactions []int64 `json:"transactions"`
}
