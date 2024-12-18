package collectors

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	cmdCollectorCreate "github.com/dpc-sdp/sumocli/pkg/cmd/collectors/create"
	cmdCollectorDelete "github.com/dpc-sdp/sumocli/pkg/cmd/collectors/delete"
	cmdCollectorGet "github.com/dpc-sdp/sumocli/pkg/cmd/collectors/get"
	cmdCollectorList "github.com/dpc-sdp/sumocli/pkg/cmd/collectors/list"
	cmdCollectorUpdate "github.com/dpc-sdp/sumocli/pkg/cmd/collectors/update"
	cmdCollectorUpgrade "github.com/dpc-sdp/sumocli/pkg/cmd/collectors/upgrade"
	"github.com/spf13/cobra"
)

func NewCmdCollectors(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collectors <command>",
		Short: "Manages collectors",
	}

	cmd.AddCommand(cmdCollectorCreate.NewCmdCollectorCreate(client))
	cmd.AddCommand(cmdCollectorDelete.NewCmdCollectorDelete(client))
	cmd.AddCommand(cmdCollectorGet.NewCmdCollectorGet(client))
	cmd.AddCommand(cmdCollectorList.NewCmdCollectorList(client))
	cmd.AddCommand(cmdCollectorUpdate.NewCmdCollectorUpdate(client))
	cmd.AddCommand(cmdCollectorUpgrade.NewCmdUpgradeCollectors(client))
	return cmd
}
