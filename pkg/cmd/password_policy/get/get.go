package get

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdPasswordPolicyGet(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get the current password policy.",
		Run: func(cmd *cobra.Command, args []string) {
			getPasswordPolicy(client)
		},
	}
	return cmd
}

func getPasswordPolicy(client *cip.APIClient) {
	data, response, err := client.GetPasswordPolicy()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
