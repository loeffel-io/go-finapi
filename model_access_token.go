/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// User access token info
type AccessToken struct {
	// Access token. Token has a length of up to 128 characters.
	AccessToken string `json:"access_token"`
	// Refresh token. Only set in case of grant_type='password'. Token has a length of up to 128 characters.
	RefreshToken string `json:"refresh_token,omitempty"`
	// Expiration time in seconds. A value of 0 means that the token never expires (unless it is explicitly invalidated, e.g. by revocation, or when a user gets locked).
	ExpiresIn int32 `json:"expires_in"`
	// Token type (it's always 'bearer')
	TokenType string `json:"token_type"`
	// Requested scopes (it's always 'all')
	Scope string `json:"scope"`
}
