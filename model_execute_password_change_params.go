/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Execute password change parameters
type ExecutePasswordChangeParams struct {
	// User identifier
	UserId string `json:"userId"`
	// User's new password. Minimum length is 6, and maximum length is 36.
	Password string `json:"password"`
	// Decrypted password change token (the token that you received from the /users/requestPasswordChange service, decrypted with your data decryption key)
	PasswordChangeToken string `json:"passwordChangeToken"`
}