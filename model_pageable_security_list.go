/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for page of securities
type PageableSecurityList struct {
	// List of securities
	Securities []Security `json:"securities"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}