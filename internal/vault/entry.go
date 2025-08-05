package vault

import "time"

// EntryType constants
const (
	TypeAccount = "account"
	TypeAPIKey  = "api_key"
	TypeSSH     = "ssh"
	TypeGPG     = "gpg"
	TypeNote    = "note"
)

// Entry represents a stored secret
type Entry struct {
	ID        string    `json:"id"`
	Label     string    `json:"label"`
	Type      string    `json:"type"`
	Created   time.Time `json:"created"`
	Encrypted string    `json:"encrypted"`
}
