package cmd

import (
	"fmt"
	"os/user"

	"github.com/spf13/cobra"

	"github.com/Clairvoy/internal/vault"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register [username]",
	Short: "Register a new user vault",
	Long: `Creates a new vault namespace for the specified user.
The vault directory ~/.vault/<username>/ and its subfolders will be initialized.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userName := args[0]

		if _, err := user.Lookup(userName); err != nil {
			fmt.Printf("Error: system user '%s' not found", userName)
			return
		}

		if err := vault.RegisterUser(userName); err != nil {
			fmt.Printf("Error registering user vault: %v", err)
			return
		}
		fmt.Printf("User vault created at ~/.vault/%s/", userName)
	},
}
	
func init() {
	rootCmd.AddCommand(registerCmd)
}
