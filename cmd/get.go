package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Clairvoy/internal/cli"
	"github.com/Clairvoy/internal/vault"
	"github.com/spf13/cobra"
	"os/user"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [id|label]",
	Short: "Retrieve a secret",
	Args:  cobra.ExactArgs(1),
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
		entry, plain, err := vault.GetEntry(base, salt, master, args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("=== %s (%s) ===", entry.Label, entry.Type)
		switch entry.Type {
		case vault.TypeAccount, vault.TypeAPIKey:
			var obj map[string]string
			json.Unmarshal(plain, &obj)
			for k, v := range obj {
				fmt.Printf("%s: %s", k, v)
			}
		default:
			fmt.Println(string(plain))
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
