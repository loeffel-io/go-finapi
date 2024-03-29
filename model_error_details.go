/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Error details
type ErrorDetails struct {
	// Error message
	Message string `json:"message,omitempty"`
	// Error code. See the documentation of the individual services for details about what values may be returned.
	Code string `json:"code,omitempty"`
	// Error type. BUSINESS errors depict German error messages for the user, e.g. from a bank server. TECHNICAL errors depict internal errors.
	Type_ string `json:"type,omitempty"`
}
