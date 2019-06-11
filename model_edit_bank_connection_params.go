/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for bank connection edit params
type EditBankConnectionParams struct {
	// New online banking user ID. If you do not want to change the current user ID let this field remain unset. In case you need to use finAPI's web form to let the user update the field, just set the field to any value, so that the service recognizes that you wish to use the web form flow. Note that you cannot clear the current user ID, i.e. a bank connection must always have a user ID (except for when it is a 'demo connection'). Max length: 64.
	BankingUserId string `json:"bankingUserId,omitempty"`
	// New online banking customer ID. If you do not want to change the current customer ID let this field remain unset. In case you need to use finAPI's web form to let the user update the field, just set the field to non-empty value, so that the service recognizes that you wish to use the web form flow. If you want to clear the current customer ID, set the field's value to an empty string (\"\"). Max length: 64.
	BankingCustomerId string `json:"bankingCustomerId,omitempty"`
	// New online banking PIN. If you do not want to change the current PIN let this field remain unset. In case you need to use finAPI's web form to let the user update the field, just set the field to non-empty value, so that the service recognizes that you wish to use the web form flow. If you want to clear the current PIN, set the field's value to an empty string (\"\").<br/><br/>NOTE: Before you set this field, please regard the 'pinsAreVolatile' flag of this connection's bank.<br/>Any symbols are allowed. Max length: 170.
	BankingPin string `json:"bankingPin,omitempty"`
	// New default two-step-procedure. Must match the 'procedureId' of one of the procedures that are listed in the bank connection. If you do not want to change this field let it remain unset. If you want to clear the current default two-step-procedure, set the field's value to an empty string (\"\").
	DefaultTwoStepProcedureId string `json:"defaultTwoStepProcedureId,omitempty"`
	// New name for the bank connection. Maximum length is 64. If you do not want to change the current name let this field remain unset. If you want to clear the current name, set the field's value to an empty string (\"\").
	Name string `json:"name,omitempty"`
}