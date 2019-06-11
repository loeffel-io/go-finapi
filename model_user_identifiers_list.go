/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for list of identifiers of deleted users, and not deleted users (in ascending order)
type UserIdentifiersList struct {
	// List of identifiers of deleted users (in ascending order)
	DeletedUsers []string `json:"deletedUsers"`
	// List of identifiers of not deleted users (in ascending order)
	NonDeletedUsers []string `json:"nonDeletedUsers"`
}
