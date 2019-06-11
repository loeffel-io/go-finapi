/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for pagination information
type Paging struct {
	// Current page number
	Page int32 `json:"page"`
	// Current page size (number of entries in this page)
	PerPage int32 `json:"perPage"`
	// Total number of pages
	PageCount int32 `json:"pageCount"`
	// Total number of entries across all pages
	TotalCount int64 `json:"totalCount"`
}
