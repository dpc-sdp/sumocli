package access_keys

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	cmdAccessKeysCreate "github.com/dpc-sdp/sumocli/pkg/cmd/access_keys/create"
	cmdAccessKeysDelete "github.com/dpc-sdp/sumocli/pkg/cmd/access_keys/delete"
	cmdAccessKeysListAll "github.com/dpc-sdp/sumocli/pkg/cmd/access_keys/list_all"
	cmdAccessKeysListPersonal "github.com/dpc-sdp/sumocli/pkg/cmd/access_keys/list_personal"
	cmdAccessKeysUpdate "github.com/dpc-sdp/sumocli/pkg/cmd/access_keys/update"
	"github.com/spf13/cobra"
)

func NewCmdAccessKeys(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access-keys",
		Short: "Manage access keys",
		Long:  "Commands that allow you to manage access keys in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdAccessKeysCreate.NewCmdAccessKeysCreate(client))
	cmd.AddCommand(cmdAccessKeysDelete.NewCmdAccessKeysDelete(client))
	cmd.AddCommand(cmdAccessKeysListAll.NewCmdAccessKeysListAll(client))
	cmd.AddCommand(cmdAccessKeysListPersonal.NewCmdAccessKeysListPersonal(client))
	cmd.AddCommand(cmdAccessKeysUpdate.NewCmdAccessKeysUpdate(client))
	return cmd
}
