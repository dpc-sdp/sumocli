package service_allowlist

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	cmdServiceAllowlistAdd "github.com/dpc-sdp/sumocli/pkg/cmd/service_allowlist/add"
	cmdServiceAllowlistDisable "github.com/dpc-sdp/sumocli/pkg/cmd/service_allowlist/disable"
	cmdServiceAllowlistEnable "github.com/dpc-sdp/sumocli/pkg/cmd/service_allowlist/enable"
	cmdServiceAllowlistList "github.com/dpc-sdp/sumocli/pkg/cmd/service_allowlist/list"
	cmdServiceAllowlistRemove "github.com/dpc-sdp/sumocli/pkg/cmd/service_allowlist/remove"
	cmdServiceAllowlistStatus "github.com/dpc-sdp/sumocli/pkg/cmd/service_allowlist/status"
	"github.com/spf13/cobra"
)

func NewCmdServiceAllowlist(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service-allowlist",
		Short: "Manage the service allowlist",
		Long:  "Commands that all you to manage the Service Allowlist in your Sumo Logic tenant",
	}
	cmd.AddCommand(cmdServiceAllowlistAdd.NewCmdServiceAllowlistAdd(client))
	cmd.AddCommand(cmdServiceAllowlistDisable.NewCmdServiceAllowlistDisable(client))
	cmd.AddCommand(cmdServiceAllowlistEnable.NewCmdServiceAllowListEnable(client))
	cmd.AddCommand(cmdServiceAllowlistList.NewCmdServiceAllowlistList(client))
	cmd.AddCommand(cmdServiceAllowlistRemove.NewCmdServiceAllowlistRemove())
	cmd.AddCommand(cmdServiceAllowlistStatus.NewCmdServiceAllowlistStatus(client))
	return cmd
}
