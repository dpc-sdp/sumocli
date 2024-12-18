package get_personal_folder

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdGetPersonalFolder(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-personal-folder",
		Short: "Get the personal folder of the current user.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getPersonalFolder(client)
		},
	}
	return cmd
}

func getPersonalFolder(client *cip.APIClient) {
	data, response, err := client.GetPersonalFolder()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
