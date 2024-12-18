package partitions

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	NewCmdPartitionsCancelRetentionUpdate "github.com/dpc-sdp/sumocli/pkg/cmd/partitions/cancel_retention_update"
	NewCmdPartitionsCreate "github.com/dpc-sdp/sumocli/pkg/cmd/partitions/create"
	NewCmdPartitionsDecommission "github.com/dpc-sdp/sumocli/pkg/cmd/partitions/decommission"
	NewCmdPartitionsGet "github.com/dpc-sdp/sumocli/pkg/cmd/partitions/get"
	NewCmdPartitionsList "github.com/dpc-sdp/sumocli/pkg/cmd/partitions/list"
	NewCmdPartitionsUpdate "github.com/dpc-sdp/sumocli/pkg/cmd/partitions/update"
	"github.com/spf13/cobra"
)

func NewCmdPartitions(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "partitions",
		Short: "Manage partitions",
		Long:  "Creating a Partition allows you to improve search performance by searching over a smaller number of messages.",
	}
	cmd.AddCommand(NewCmdPartitionsCancelRetentionUpdate.NewCmdPartitionsCancelRetentionUpdate(client))
	cmd.AddCommand(NewCmdPartitionsCreate.NewCmdPartitionCreate(client))
	cmd.AddCommand(NewCmdPartitionsDecommission.NewCmdPartitionsDecommission(client))
	cmd.AddCommand(NewCmdPartitionsGet.NewCmdPartitionsGet(client))
	cmd.AddCommand(NewCmdPartitionsList.NewCmdPartitionsList(client))
	cmd.AddCommand(NewCmdPartitionsUpdate.NewCmdPartitionUpdate(client))
	return cmd
}
