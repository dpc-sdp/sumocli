package health_events

import (
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	NewCmdHealthEventsGet "github.com/dpc-sdp/sumocli/pkg/cmd/health_events/get"
	NewCmdHealthEventsList "github.com/dpc-sdp/sumocli/pkg/cmd/health_events/list"
	"github.com/spf13/cobra"
)

func NewCmdHealthEvents(client *cip.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health-events",
		Short: "Manages health events",
		Long: "Health Events allow you to keep track of the health of your Collectors and Sources. " +
			"You can use them to find and investigate common errors and warnings that are known to cause collection issues.",
	}
	cmd.AddCommand(NewCmdHealthEventsGet.NewCmdHealthEventsGet(client))
	cmd.AddCommand(NewCmdHealthEventsList.NewCmdHealthEventsList(client))
	return cmd
}
