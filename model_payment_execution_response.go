/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Bank server's response to Money Transfer / Direct Debit execution
type PaymentExecutionResponse struct {
	// Technical message from the bank server, confirming the success of the request. Typically, you would not want to present this message to the user. Note that this field may not be set. However if it is not set, it does not necessarily mean that there was an error in processing the request.
	SuccessMessage string `json:"successMessage,omitempty"`
	// In some cases, a bank server may accept the requested order, but return a warn message. This message may be of technical nature, but could also be of interest to the user.
	WarnMessage string `json:"warnMessage,omitempty"`
	// Payment identifier. Can be used to retrieve the status of the payment (see 'Get payments' service).
	PaymentId int64 `json:"paymentId"`
}
