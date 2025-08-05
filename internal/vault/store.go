package vault

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Clairvoy/internal/crypto"
	"github.com/Clairvoy/internal/fileutil"
	"github.com/google/uuid"
)

var types = []string{TypeAccount, TypeAPIKey, TypeSSH, TypeGPG, TypeNote}

const (
	rootDirName = ".vault"
	saltName    = "salt"
)

// RegisterUser initializes a vault namespace for a new user.
func RegisterUser(username string) error {
	base := userBasePath(username)
	// fail if already exists
	if _, err := os.Stat(base); !os.IsNotExist(err) {
		return fmt.Errorf("vault for user '%s' already exists", username)
	}
	// call init logic to create dir and salt
	_, _, err := initVault(username)
	return err
}

// Init ensures the user's vault exists and returns its base path and salt.
func Init(username string) (string, []byte, error) {
	base := userBasePath(username)
	saltPath := filepath.Join(base, saltName)
	// if not exist, register first
	if _, err := os.Stat(saltPath); os.IsNotExist(err) {
		if err := RegisterUser(username); err != nil {
			return "", nil, err
		}
	}
	// load salt and return
	saltData, err := os.ReadFile(saltPath)
	if err != nil {
		return "", nil, err
	}
	salt, err := base64.StdEncoding.DecodeString(string(saltData))
	if err != nil {
		return "", nil, err
	}
	return base, salt, nil
}

// initVault creates directories and writes salt, used internally.
func initVault(username string) (string, []byte, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", nil, err
	}
	base := filepath.Join(home, rootDirName, username)
	// create main and subtype dirs
	if err := os.MkdirAll(base, 0o700); err != nil {
		return "", nil, err
	}
	for _, t := range types {
		if err := os.MkdirAll(filepath.Join(base, t), 0o700); err != nil {
			return "", nil, err
		}
	}
	// generate salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	encoded := base64.StdEncoding.EncodeToString(salt)
	saltPath := filepath.Join(base, saltName)
	if err := os.WriteFile(saltPath, []byte(encoded), 0o600); err != nil {
		return "", nil, err
	}
	return base, salt, nil
}

// helper for base path
func userBasePath(username string) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, rootDirName, username)
}

// AddEntry encrypts and stores a new entry
func AddEntry(base string, salt []byte, master, typ, label string, plaintext []byte) error {
	if !validType(typ) {
		return fmt.Errorf("invalid type %s", typ)
	}
	key := crypto.DeriveKey(master, salt)
	enc, err := crypto.Encrypt(plaintext, key)
	if err != nil {
		return err
	}
	entry := Entry{uuid.NewString(), label, typ, time.Now().UTC(), enc}
	path := filepath.Join(base, typ, entry.ID+".json")
	return fileutil.AtomicWriteJSON(path, &entry, 0o600)
}

// ListEntries returns all entry metadata
func ListEntries(base string) ([]Entry, error) {
	var out []Entry
	for _, t := range types {
		dir := filepath.Join(base, t)
		files, _ := os.ReadDir(dir)
		for _, f := range files {
			data, _ := os.ReadFile(filepath.Join(dir, f.Name()))
			var e Entry
			json.Unmarshal(data, &e)
			out = append(out, e)
		}
	}
	return out, nil
}

// GetEntry finds and decrypts an entry
func GetEntry(base string, salt []byte, master, idOrLabel string) (*Entry, []byte, error) {
	for _, t := range types {
		dir := filepath.Join(base, t)
		files, _ := os.ReadDir(dir)
		for _, f := range files {
			data, _ := os.ReadFile(filepath.Join(dir, f.Name()))
			var e Entry
			json.Unmarshal(data, &e)
			if e.ID == idOrLabel || e.Label == idOrLabel {
				key := crypto.DeriveKey(master, salt)
				pt, err := crypto.Decrypt(e.Encrypted, key)
				return &e, pt, err
			}
		}
	}
	return nil, nil, errors.New("not found")
}

// RemoveEntry deletes the entry file
func RemoveEntry(base, idOrLabel string) error {
	for _, t := range types {
		dir := filepath.Join(base, t)
		files, _ := os.ReadDir(dir)
		for _, f := range files {
			data, _ := os.ReadFile(filepath.Join(dir, f.Name()))
			var e Entry
			json.Unmarshal(data, &e)
			if e.ID == idOrLabel || e.Label == idOrLabel {
				return os.Remove(filepath.Join(dir, f.Name()))
			}
		}
	}
	return errors.New("not found")
}

// validType checks if provided type is supported
func validType(t string) bool {
	for _, tt := range types {
		if tt == t {
			return true
		}
	}
	return false
}
