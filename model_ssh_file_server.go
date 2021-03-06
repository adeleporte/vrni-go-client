/*
 * vRealize Network Insight API Reference
 *
 * vRealize Network Insight API Reference
 *
 * API version: 1.1.8
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vrni
// SshFileServer struct for SshFileServer
type SshFileServer struct {
	// IP address or hostname of destination SSH file server
	ServerAddress string `json:"server_address,omitempty"`
	// File transfer port
	Port int32 `json:"port,omitempty"`
	// username to login ssh server
	Username string `json:"username,omitempty"`
	// password for the user to login ssh server
	Password string `json:"password,omitempty"`
	// directory on file server to store/fetch backups
	BackupDirectory string `json:"backup_directory,omitempty"`
	// filename of the backup to be restored (used during restore only)
	BackupFileName string `json:"backup_file_name,omitempty"`
}
