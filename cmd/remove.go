package cmd

import (
	"fmt"
	"github.com/Clairvoy/internal/vault"
	"os/user"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [id|label]",
	Short: "Remove a secret",
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
		base, _, err := vault.Init(usr.Username)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := vault.RemoveEntry(base, args[0]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Removed.")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
