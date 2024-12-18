package update_subdomain

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdAccountUpdateSubdomain(client *cip.APIClient) *cobra.Command {
	var subdomain string
	cmd := &cobra.Command{
		Use:   "update-subdomain",
		Short: "Update a subdomain. Only the Account Owner can update the subdomain.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			updateSubdomain(subdomain, client)
		},
	}
	cmd.Flags().StringVar(&subdomain, "subdomain", "", "Specify a new subdomain (minimum 4 and maximum 63 characters)")
	cmd.MarkFlagRequired("subdomain")
	return cmd
}

func updateSubdomain(subdomain string, client *cip.APIClient) {
	data, response, err := client.UpdateSubdomain(types.ConfigureSubdomainRequest{
		Subdomain: subdomain,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
