package get

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdTokensGet(client *cip.APIClient) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a token with the given identifier in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getToken(id, client)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the token to retrieve")
	cmd.MarkFlagRequired("id")
	return cmd
}

func getToken(id string, client *cip.APIClient) {
	data, response, err := client.GetToken(id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
