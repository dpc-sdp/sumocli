package users

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	cmdUserChange "github.com/dpc-sdp/sumocli/pkg/cmd/users/change_email"
	cmdUserCreate "github.com/dpc-sdp/sumocli/pkg/cmd/users/create"
	cmdUserDelete "github.com/dpc-sdp/sumocli/pkg/cmd/users/delete"
	cmdUserDisable "github.com/dpc-sdp/sumocli/pkg/cmd/users/disable_mfa"
	cmdUserGet "github.com/dpc-sdp/sumocli/pkg/cmd/users/get"
	cmdUserList "github.com/dpc-sdp/sumocli/pkg/cmd/users/list"
	cmduserReset "github.com/dpc-sdp/sumocli/pkg/cmd/users/reset_password"
	cmdUserUnlock "github.com/dpc-sdp/sumocli/pkg/cmd/users/unlock"
	cmdUserUpdate "github.com/dpc-sdp/sumocli/pkg/cmd/users/update"
	"github.com/spf13/cobra"
)

func NewCmdUser(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users <command>",
		Short: "Manage users",
		Long:  "Work with Sumo Logic users",
	}

	cmd.AddCommand(cmdUserChange.NewCmdUserChangeEmail(client))
	cmd.AddCommand(cmdUserCreate.NewCmdUserCreate(client))
	cmd.AddCommand(cmdUserDelete.NewCmdUserDelete(client))
	cmd.AddCommand(cmdUserDisable.NewCmdUserDisableMFA(client))
	cmd.AddCommand(cmdUserGet.NewCmdGetUser(client))
	cmd.AddCommand(cmdUserList.NewCmdUserList(client))
	cmd.AddCommand(cmduserReset.NewCmdUserResetPassword(client))
	cmd.AddCommand(cmdUserUnlock.NewCmdUnlockUser(client))
	cmd.AddCommand(cmdUserUpdate.NewCmdUserUpdate(client))
	return cmd
}
