/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.67.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Data of notification rule
type NotificationRule struct {
	// Notification rule identifier
	Id int64 `json:"id"`
	// Trigger event type
	TriggerEvent string `json:"triggerEvent"`
	// Additional parameters that are specific to the trigger event type. Please refer to the documentation for details.
	Params map[string]string `json:"params,omitempty"`
	// The string that finAPI includes into the notifications that it sends based on this rule.
	CallbackHandle string `json:"callbackHandle,omitempty"`
	// Whether the notification messages that will be sent based on this rule contain encrypted detailed data or not
	IncludeDetails bool `json:"includeDetails"`
}