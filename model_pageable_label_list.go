/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Label resources with paging information
type PageableLabelList struct {
	// Data of labels
	Labels []Label `json:"labels"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}
