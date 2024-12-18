package recover_subdomain

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdAccountRecoverSubdomain(client *cip.APIClient) *cobra.Command {
	var email string
	cmd := &cobra.Command{
		Use:   "recover-subdomain",
		Short: "Send an email with the subdomain information for a user with the given email address.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			recoverSubdomain(email, client)
		},
	}
	cmd.Flags().StringVar(&email, "email", "", "Specify an email address of the user to get subdomain information")
	cmd.MarkFlagRequired("email")
	return cmd
}

func recoverSubdomain(email string, client *cip.APIClient) {
	response, err := client.RecoverSubdomains(email)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(nil, response, err, "An email containing information about associated subdomains for the given email was sent.")
	}
}
