package cmd

import (
	"fmt"
	"github.com/Clairvoy/internal/vault"
	"os/user"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secrets",
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
		base, _, err := vault.Init(usr.Username)
		if err != nil {
			fmt.Println(err)
			return
		}
		entries, err := vault.ListEntries(base)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, e := range entries {
			fmt.Printf("- %s (%s) added %s id=%s\n",
				e.Label, e.Type, e.Created.Format(time.RFC3339), e.ID)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
