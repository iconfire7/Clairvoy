/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os/user"

	"github.com/spf13/cobra"

	"github.com/Clairvoy/internal/cli"
	"github.com/Clairvoy/internal/vault"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new secret",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			fmt.Println("Failed to get current user:", err)
			return
		}

		if _, err := user.Lookup(usr.Username); err != nil {
			fmt.Printf("Error: system user '%s' not found", usr.Username)
			return
		}

		base, salt, err := vault.Init(usr.Username)
		if err != nil {
			fmt.Println(err)
			return
		}

		master, _ := cli.PromptSecret("Master password: ")

		fmt.Println("Types: account, api_key, ssh, gpg, note")
		typ := cli.PromptLine("Type: ")
		label := cli.PromptLine("Label: ")
		var payload []byte

		switch typ {
		case vault.TypeAccount:
			login := cli.PromptLine("Login: ")
			pass, _ := cli.PromptSecret("Password: ")
			obj := map[string]string{"login": login, "password": pass}
			payload, _ = json.Marshal(obj)
		case vault.TypeAPIKey:
			service := cli.PromptLine("Service: ")
			key := cli.PromptLine("API Key: ")
			obj := map[string]string{"service": service, "key": key}
			payload, _ = json.Marshal(obj)
		case vault.TypeSSH:
			raw := cli.PromptMultiline("Paste SSH private key")
			payload = []byte(raw)
		case vault.TypeGPG:
			raw := cli.PromptMultiline("Paste GPG private key")
			payload = []byte(raw)
		case vault.TypeNote:
			note := cli.PromptMultiline("Enter note")
			payload = []byte(note)
		default:
			fmt.Println("Unknown type")
			return
		}
		if err := vault.AddEntry(base, salt, master, typ, label, payload); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Added.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
