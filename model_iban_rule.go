/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for an IBAN rule
type IbanRule struct {
	// Rule identifier
	Id int64 `json:"id"`
	// The category that this rule assigns to the transactions that it matches
	Category *Category `json:"category"`
	// Direction for the rule. 'Income' means that the rule applies to transactions with a positive amount only, 'Spending' means it applies to transactions with a negative amount only.
	Direction string `json:"direction"`
	// Timestamp of when the rule was created, in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time)
	CreationDate string `json:"creationDate"`
	// The IBAN for this rule
	Iban string `json:"iban"`
}
