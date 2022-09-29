/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateGroup struct {
	// Group name
	GroupName string `json:"groupName"`
	// Members
	Members []string `json:"members,omitempty"`
	// MetaData
	MetaData []map[string]string `json:"metaData,omitempty"`
}
