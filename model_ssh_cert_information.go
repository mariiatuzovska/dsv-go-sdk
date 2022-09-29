/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SshCertInformation struct {
	// Path to secret containing leaf certificate
	LeafCAPath string `json:"leafCAPath"`
	// A list of principals on a certificate (user or host names)
	Principals []string `json:"principals"`
	// Path to secret containing root certificate
	RootCAPath string `json:"rootCAPath"`
	// TTL for a generated certificate (in hours)
	Ttl int64 `json:"ttl"`
}
