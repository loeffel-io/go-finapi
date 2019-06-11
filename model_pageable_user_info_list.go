/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for users information with paging info
type PageableUserInfoList struct {
	// List of users information
	Users []UserInfo `json:"users"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}
