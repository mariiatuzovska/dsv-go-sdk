/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Creation properties of a user
type DsvUserCreateModel struct {
	// The display name of the user
	DisplayName string `json:"displayName,omitempty"`
	// External identifier, such as an AWS arn for 3rd party authentication
	ExternalId string `json:"externalId,omitempty"`
	// User's password (not required if using 3rd party auth)
	Password string `json:"password,omitempty"`
	// Provider name defined in the authentication settings section of configuration
	Provider string `json:"provider,omitempty"`
	// The name of the user
	UserName string `json:"userName"`
}
