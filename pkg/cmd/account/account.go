package account

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	NewCmdAccountCreateSubdomain "github.com/dpc-sdp/sumocli/pkg/cmd/account/create-subdomain"
	NewCmdAccountDeleteSubdomain "github.com/dpc-sdp/sumocli/pkg/cmd/account/delete-subdomain"
	NewCmdAccountGetOwner "github.com/dpc-sdp/sumocli/pkg/cmd/account/get-owner"
	NewCmdAccountGetSubdomain "github.com/dpc-sdp/sumocli/pkg/cmd/account/get-subdomain"
	NewCmdAccountRecoverSubdomain "github.com/dpc-sdp/sumocli/pkg/cmd/account/recover-subdomain"
	NewCmdAccountUpdateSubdomain "github.com/dpc-sdp/sumocli/pkg/cmd/account/update-subdomain"
	"github.com/spf13/cobra"
)

func NewCmdAccount(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account",
		Short: "Manage account settings",
		Long:  "Commands that allow you to manage your Sumo Logic account settings",
	}
	cmd.AddCommand(NewCmdAccountCreateSubdomain.NewCmdAccountCreateSubdomain(client))
	cmd.AddCommand(NewCmdAccountDeleteSubdomain.NewCmdAccountDeleteSubdomain(client))
	cmd.AddCommand(NewCmdAccountGetOwner.NewCmdAccountGetOwner(client))
	cmd.AddCommand(NewCmdAccountGetSubdomain.NewCmdAccountGetSubdomain(client))
	cmd.AddCommand(NewCmdAccountRecoverSubdomain.NewCmdAccountRecoverSubdomain(client))
	cmd.AddCommand(NewCmdAccountUpdateSubdomain.NewCmdAccountUpdateSubdomain(client))
	return cmd
}
