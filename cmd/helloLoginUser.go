package cmd

import (
	"dev-docs-cli/pkg/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var sayHelloLoginUser = &cobra.Command{
	Use:   "hello",
	Short: "Greet the logged-in user",
	Long:  `This command greets the logged-in user by reading the user information from a token file.`,
	Run: func(cmd *cobra.Command, args []string) {
		exePath, err := os.Executable()
		if err != nil {
			fmt.Println("Error getting executable path:", err)
			return
		}

		rootPath := filepath.Dir(exePath)
		tokenPath := filepath.Join(rootPath, "devdocs_token")

		user, err := utils.ReadTokenFromFile(tokenPath)
		if err != nil {
			fmt.Println("User is not Authenticated")
			return
		}

		tokenExpired, err := utils.IsTokenExpired(user.AccessToken)
		if err != nil {
			fmt.Println("Error checking token expiration:", err)
			return
		}

		if tokenExpired {
			fmt.Println("Token expired. Please authenticate first.")
			fmt.Println("Use this command for Login: browser-login")
		} else {
			fmt.Printf("Hello, %s! Your email is %s and your user ID is %s.\n", user.Username, user.Email, user.UserID)
		}
	},
}

func init() {
	rootCmd.AddCommand(sayHelloLoginUser)
}
