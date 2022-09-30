/*
 * DevOps Secrets Vault API
 *
 * The purpose of this application is to provide a simple service for storing and getting secrets
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DsvSiemResponse struct {
	// Authentication token
	Auth string `json:"auth,omitempty"`
	// Authentication method (token)
	AuthMethod string `json:"authMethod,omitempty"`
	// Endpoint
	Endpoint string `json:"endpoint,omitempty"`
	// Failed is true if send has failed too many times, false otherwise
	Failed bool `json:"failed,omitempty"`
	// Number of failed send events
	FailedEvents int64 `json:"failedEvents,omitempty"`
	// Collect Server IP/FQDN
	Host string `json:"host,omitempty"`
	// The unique id for this item
	Id string `json:"id,omitempty"`
	// Logging format (e.g. \"rfc5424\" for syslog)
	LoggingFormat string `json:"loggingFormat,omitempty"`
	// Name of registered SIEM endpoint, similar to path
	Name string `json:"name,omitempty"`
	// Engine pool name, used when sending request to a DSV engine instance
	Pool string `json:"pool,omitempty"`
	// Collect Server Port
	Port int64 `json:"port,omitempty"`
	// Type of protocol (\"tcp\", \"udp\", \"http\", \"https\", \"tls\")
	Protocol string `json:"protocol,omitempty"`
	// Denotes whether the endpoint should be accessed through a DSV engine instance
	SendToEngine bool `json:"sendToEngine,omitempty"`
	// Type of endpoint (\"syslog\", \"cef\", \"json\", \"splunk\")
	SiemType string `json:"siemType,omitempty"`
}
