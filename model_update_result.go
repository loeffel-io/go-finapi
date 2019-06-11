/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for a status of bank connection update
type UpdateResult struct {
	// Note that 'OK' just means that finAPI could successfully connect and log in to the bank server. However, it does not necessarily mean that all accounts could be updated successfully. For the latter, please refer to the 'status' field of the Account resource.
	Result string `json:"result"`
	// In case the update result is not <code>OK</code>, this field may contain an error message with details about why the update failed (it is not guaranteed that a message is available though). In case the update result is <code>OK</code>, the field will always be null.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// In case the update result is not <code>OK</code>, this field contains the type of the error that occurred. BUSINESS means that the bank server responded with a non-technical error message for the user. TECHNICAL means that some internal error has occurred in finAPI or at the bank server.
	ErrorType string `json:"errorType,omitempty"`
	// Time of the update. The value is returned in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	Timestamp string `json:"timestamp"`
}