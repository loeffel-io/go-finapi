/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Parameters of keyword rule
type KeywordRuleParams struct {
	// ID of the category that this rule should assign to the matching transactions
	CategoryId int64 `json:"categoryId"`
	// Direction for the rule. 'Income' means that the rule applies to transactions with a positive amount only, 'Spending' means it applies to transactions with a negative amount only. 'Both' means that it applies to both kind of transactions. Note that in case of 'Both', finAPI will create two individual rules (one with direction 'Income' and one with direction 'Spending').
	Direction string `json:"direction"`
	// Set of keywords for the rule (Keywords are regarded case-insensitive). The minimum number of keywords is 1. The maximum number of keywords is 100.
	Keywords []string `json:"keywords"`
}
