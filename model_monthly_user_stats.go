/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Information about a user's data or activities for a certain month
type MonthlyUserStats struct {
	// The month that the contained information applies to, in the format 'YYYY-MM'.
	Month string `json:"month"`
	// Minimum count of bank connections that this user has had at any point during the month.
	MinBankConnectionCount int32 `json:"minBankConnectionCount"`
	// Maximum count of bank connections that this user has had at any point during the month.
	MaxBankConnectionCount int32 `json:"maxBankConnectionCount"`
}
