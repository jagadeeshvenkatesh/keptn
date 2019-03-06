package cmd

import (
	"errors"
	"fmt"

	"github.com/keptn/keptn/cli/utils"
	"github.com/keptn/keptn/cli/utils/credentialmanager"
	"github.com/knative/pkg/cloudevents"
	"github.com/spf13/cobra"
)

type configData struct {
	Org   *string `json:"org"`
	User  *string `json:"user"`
	Token *string `json:"token"`
}

var config configData

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configures the GitHub org, user and token in the keptn installation.",
	Long: `Configures the GitHub Organization, the GitHub user, and the GitHub
	token in the keptn installation. Usage of \"configure\":

keptn configure --org=MyOrg --user=keptnUser --token=XYZ`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting to configure Github org, user and token")

		builder := cloudevents.Builder{
			Source:    "https://github.com/keptn/keptn/cli#configure",
			EventType: "configure",
		}
		endPoint, apiToken, err := credentialmanager.GetCreds()
		if err != nil || endPoint == "" {
			utils.Info.Printf("Configure called without beeing authenticated.")
			return errors.New("This command requires to be authenticated. See \"keptn auth\" for details")
		}
		configureEndPoint := endPoint + "config"
		err = utils.Send(configureEndPoint, apiToken, builder, config)
		if err != nil {
			utils.Error.Printf("Configure command was unsuccessful. Details: %v", err)
			return err
		}
		fmt.Println("Successfully configured Github org, user and token in the keptn installation.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	config.Org = configureCmd.Flags().StringP("org", "o", "", "The GitHub organization")
	configureCmd.MarkFlagRequired("org")
	config.User = configureCmd.Flags().StringP("user", "u", "", "The GitHub user")
	configureCmd.MarkFlagRequired("user")
	config.Token = configureCmd.Flags().StringP("token", "t", "", "The GitHub token")
	configureCmd.MarkFlagRequired("token")
}