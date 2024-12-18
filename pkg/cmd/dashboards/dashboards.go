package dashboards

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	NewCmdDashboardCreate "github.com/dpc-sdp/sumocli/pkg/cmd/dashboards/create"
	NewCmdDashboardDelete "github.com/dpc-sdp/sumocli/pkg/cmd/dashboards/delete"
	NewCmdDashboardsGet "github.com/dpc-sdp/sumocli/pkg/cmd/dashboards/get"
	NewCmdDashboardsUpdate "github.com/dpc-sdp/sumocli/pkg/cmd/dashboards/update"
	"github.com/spf13/cobra"
)

func NewCmdDashboards(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboards",
		Short: "Manage dashboards (New)",
		Long:  "Commands that allow you to create, modify or delete new dashboards.",
	}
	cmd.AddCommand(NewCmdDashboardCreate.NewCmdDashboardsCreate(client))
	cmd.AddCommand(NewCmdDashboardDelete.NewCmdDashboardsDelete(client))
	cmd.AddCommand(NewCmdDashboardsGet.NewCmdDashboardsGet(client))
	cmd.AddCommand(NewCmdDashboardsUpdate.NewCmdDashboardsUpdate(client))
	return cmd
}
