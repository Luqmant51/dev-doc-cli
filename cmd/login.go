package cmd

import (
	"dev-docs-cli/pkg/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var login = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with Keycloak using the browser and save token",
	Long:  `Opens a browser to authenticate with Keycloak and saves the access token in a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		keycloakURL := viper.GetString("keycloak.url")
		realm := viper.GetString("keycloak.realm")
		clientID := viper.GetString("keycloak.client_id")
		clientSecret := viper.GetString("keycloak.client_secret")

		// Step 1: Get the device authorization endpoint
		deviceAuthURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/auth/device", keycloakURL, realm)
		request := gorequest.New()
		var deviceAuthResp map[string]interface{}
		_, _, errs := request.Post(deviceAuthURL).
			Type("form").
			Send(fmt.Sprintf("client_id=%s", clientID)).
			Send(fmt.Sprintf("client_secret=%s", clientSecret)).
			EndStruct(&deviceAuthResp)

		if len(errs) > 0 {
			log.Fatalf("Failed to initiate device authorization: %v", errs)
		}

		userCode, ok := deviceAuthResp["user_code"].(string)
		if !ok {
			log.Fatalf("Failed to get user_code from response: %v", deviceAuthResp)
		}

		verificationURI, ok := deviceAuthResp["verification_uri_complete"].(string)
		if !ok {
			log.Fatalf("Failed to get verification_uri_complete from response: %v", deviceAuthResp)
		}

		deviceCode, ok := deviceAuthResp["device_code"].(string)
		if !ok {
			log.Fatalf("Failed to get device_code from response: %v", deviceAuthResp)
		}

		interval, ok := deviceAuthResp["interval"].(float64)
		if !ok {
			log.Fatalf("Failed to get interval from response: %v", deviceAuthResp)
		}

		fmt.Printf("Please go to %s and enter the code: %s\n", verificationURI, userCode)
		browser.OpenURL(verificationURI)

		// Step 2: Poll the token endpoint
		tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", keycloakURL, realm)
		for {
			var tokenResp map[string]interface{}
			_, _, errs := request.Post(tokenURL).
				Type("form").
				Send(fmt.Sprintf("client_id=%s", clientID)).
				Send(fmt.Sprintf("client_secret=%s", clientSecret)).
				Send(fmt.Sprintf("grant_type=urn:ietf:params:oauth:grant-type:device_code")).
				Send(fmt.Sprintf("device_code=%s", deviceCode)).
				EndStruct(&tokenResp)

			if len(errs) > 0 {
				log.Fatalf("Failed to poll token endpoint: %v", errs)
			}

			if accessToken, ok := tokenResp["access_token"].(string); ok {
				// Decode the access token to get the additional information
				claims := utils.NewDecodeToken(accessToken)

				username, _ := claims["preferred_username"].(string)
				email, _ := claims["email"].(string)
				userid, _ := claims["sub"].(string)

				// Save the token and additional information to a file
				tokenFile := filepath.Join(os.Getenv("HOME"), "devdocs_token")

				// Ensure the directory exists
				err := os.MkdirAll(filepath.Dir(tokenFile), os.ModePerm)
				if err != nil {
					log.Fatalf("Failed to create directory: %v", err)
				}

				tokenData := fmt.Sprintf("AccessToken: %s\nUsername: %s\nEmail: %s\nUserID: %s", accessToken, username, email, userid)
				err = os.WriteFile(tokenFile, []byte(tokenData), 0600)
				if err != nil {
					log.Fatalf("Failed to save token: %v", err)
				}

				fmt.Println("Successfully authenticated and saved token.")
				return
			}

			if error, ok := tokenResp["error"].(string); ok {
				if error == "authorization_pending" {
					// Wait and try again
					time.Sleep(time.Duration(interval) * time.Second)
				} else {
					errorDescription, _ := tokenResp["error_description"].(string)
					log.Fatalf("Error during token polling: %v - %v", error, errorDescription)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(login)
}
