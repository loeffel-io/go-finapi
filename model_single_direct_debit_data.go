/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Data for a single direct debit
type SingleDirectDebitData struct {
	// Name of the debitor. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the debitor name. Even if the debitor name does not depict the actual registered account holder of the specified debitor account, the direct debit request might still be successful.
	DebitorName string `json:"debitorName"`
	// IBAN of the debitor's account
	DebitorIban string `json:"debitorIban"`
	// BIC of the debitor's account. Note: This field is optional if - and only if - the bank connection of the account that you want to transfer money to supports the IBAN-Only direct debit. You can find this out via GET /bankConnections/<id>. Also note that when a BIC is given, then this BIC will be used for the direct debit request independent of whether it is required or not.
	DebitorBic string `json:"debitorBic,omitempty"`
	// The amount to transfer. Must be a positive decimal number with at most two decimal places (e.g. 99.99)
	Amount float32 `json:"amount"`
	// The purpose of the transfer transaction
	Purpose string `json:"purpose,omitempty"`
	// SEPA purpose code, according to ISO 20022, external codes set.
	SepaPurposeCode string `json:"sepaPurposeCode,omitempty"`
	// Mandate ID that this direct debit order is based on.
	MandateId string `json:"mandateId"`
	// Date of the mandate that this direct debit order is based on, in the format 'YYYY-MM-DD'
	MandateDate string `json:"mandateDate"`
	// Creditor ID of the source account's holder
	CreditorId string `json:"creditorId,omitempty"`
	// End-To-End ID for the transfer transaction
	EndToEndId string `json:"endToEndId,omitempty"`
}
