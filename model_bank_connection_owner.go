/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container for a bank connection owner's data
type BankConnectionOwner struct {
	// First name
	FirstName string `json:"firstName,omitempty"`
	// Last name
	LastName string `json:"lastName,omitempty"`
	// Salutation
	Salutation string `json:"salutation,omitempty"`
	// Title
	Title string `json:"title,omitempty"`
	// Email
	Email string `json:"email,omitempty"`
	// Date of birth (in format 'YYYY-MM-DD')
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	// Post code
	PostCode string `json:"postCode,omitempty"`
	// Country
	Country string `json:"country,omitempty"`
	// City
	City string `json:"city,omitempty"`
	// Street
	Street string `json:"street,omitempty"`
	// House number
	HouseNumber string `json:"houseNumber,omitempty"`
}
