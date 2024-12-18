package update

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCmdTokensUpdate(client *cip.APIClient) *cobra.Command {
	var (
		description string
		id          string
		inactive    bool
		name        string
		version     int64
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a token with the given identifier in the token library.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			updateToken(description, id, inactive, name, version, client)
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "Specify a description for the token")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the token to update")
	cmd.Flags().BoolVar(&inactive, "inactive", false, "Set to true if you want the token to be inactive")
	cmd.Flags().StringVar(&name, "name", "", "Specify a name for the token")
	cmd.Flags().Int64Var(&version, "version", 0, "Specify a version of the token (can be retrieved by running sumocli tokens list)")
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("version")
	return cmd
}

func updateToken(description string, id string, inactive bool, name string, version int64, client *cip.APIClient) {
	var options types.TokenBaseDefinitionUpdate
	if inactive == true {
		options.Status = "Inactive"
	} else {
		options.Status = "Active"
	}
	options.Name = name
	options.Description = description
	options.Type_ = "CollectorRegistration"
	options.Version = version
	data, response, err := client.UpdateToken(options, id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
