/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for categories with paging info
type PageableCategoryList struct {
	// Categories
	Categories []Category `json:"categories"`
	// Information for pagination
	Paging *Paging `json:"paging"`
}
