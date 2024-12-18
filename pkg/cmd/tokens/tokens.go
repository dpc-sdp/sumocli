package tokens

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	cmdTokensCreate "github.com/dpc-sdp/sumocli/pkg/cmd/tokens/create"
	cmdTokensDelete "github.com/dpc-sdp/sumocli/pkg/cmd/tokens/delete"
	cmdTokensGet "github.com/dpc-sdp/sumocli/pkg/cmd/tokens/get"
	cmdTokensList "github.com/dpc-sdp/sumocli/pkg/cmd/tokens/list"
	cmdTokensUpdate "github.com/dpc-sdp/sumocli/pkg/cmd/tokens/update"
	"github.com/spf13/cobra"
)

func NewCmdTokens(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "Manage tokens",
		Long:  "Commands that allow you to manage Tokens in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdTokensCreate.NewCmdTokensCreate(client))
	cmd.AddCommand(cmdTokensDelete.NewCmdTokensDelete(client))
	cmd.AddCommand(cmdTokensGet.NewCmdTokensGet(client))
	cmd.AddCommand(cmdTokensList.NewCmdTokensList(client))
	cmd.AddCommand(cmdTokensUpdate.NewCmdTokensUpdate(client))
	return cmd
}
