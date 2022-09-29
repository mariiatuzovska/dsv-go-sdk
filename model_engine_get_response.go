/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Engine is the stored record
type EngineGetResponse struct {
	// Created date
	Created string `json:"created,omitempty"`
	// Who created the item
	CreatedBy string `json:"createdBy,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	// the id for this item
	Id string `json:"id,omitempty"`
	LastHeartbeat string `json:"lastHeartbeat,omitempty"`
	// Last updated date
	LastModified string `json:"lastModified,omitempty"`
	// Who performed the last modification
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
	Name string `json:"name,omitempty"`
	PoolName string `json:"poolName,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	// Current version
	Version string `json:"version,omitempty"`
}
