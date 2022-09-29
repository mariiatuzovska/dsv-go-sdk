/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// AddToGroupsRequest to add user to a list of groups
type AddToGroupsRequest struct {
	// List of group names
	GroupNames []string `json:"groupNames,omitempty"`
}
