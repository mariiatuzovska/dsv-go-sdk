/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Audit model
type DsvAudit struct {
	// Action performed
	Action string `json:"action,omitempty"`
	// Audit created date
	Created time.Time `json:"created,omitempty"`
	// Audit id
	Id string `json:"id,omitempty"`
	// IP Address logged from client
	Ipaddress string `json:"ipaddress,omitempty"`
	// Message additional details
	Message string `json:"message,omitempty"`
	// Resource path action performed on
	Path string `json:"path,omitempty"`
	// Security principal that performed action
	Principal string `json:"principal,omitempty"`
	// Principal item ID
	PrincipalItemId string `json:"principalItemId,omitempty"`
	// Http status code
	Status int64 `json:"status,omitempty"`
	// Tenant ID
	Tenant string `json:"tenant,omitempty"`
	// Tenant Name
	TenantName string `json:"tenantName,omitempty"`
}
