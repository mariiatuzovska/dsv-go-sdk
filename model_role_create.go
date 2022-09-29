/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RoleCreate struct {
	// Role description
	Description string `json:"description,omitempty"`
	// External identifier, such as an AWS arn for 3rd party authentication
	ExternalId string `json:"externalId,omitempty"`
	// Name of role
	Name string `json:"name"`
	// Provider name defined in the authentication settings section of configuration
	Provider string `json:"provider,omitempty"`
}
