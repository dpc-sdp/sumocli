package update

import (
	"encoding/json"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdDashboardsUpdate(client *cip.APIClient) *cobra.Command {
	var (
		id   string
		file string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a dashboard by the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			updateDashboard(file, id, client)
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "Specify the full file path to a json file containing a dashboard definition."+
		"The definition can be retrieved from running sumocli dashboards get or from exporting the dashboard in the UI."+
		"If you have exported the dashboard definition you may need to modify the panel ids before updating.")
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the dashboard to update")
	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("id")
	return cmd
}

func updateDashboard(file string, id string, client *cip.APIClient) {
	var dashboardDefinition types.DashboardRequest
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Error().Err(err).Msg("failed to read file " + file)
	}
	err = json.Unmarshal(fileData, &dashboardDefinition)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal file")
	}
	data, response, err := client.UpdateDashboard(dashboardDefinition, id)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
