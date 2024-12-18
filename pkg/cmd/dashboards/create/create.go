package create

import (
	"encoding/json"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/dpc-sdp/sumocli/internal/authentication"
	"github.com/dpc-sdp/sumocli/pkg/cmdutils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdDashboardsCreate(client *cip.APIClient) *cobra.Command {
	var file string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new dashboard.",
		Long: "Note: When exporting a dashboard spec from the Sumo Logic portal ensure that you have the timeRange.to object set as well as the " +
			"timeRange.from set otherwise you will get errors when trying to create.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			createDashboard(file, client)
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "Specify the full file path to a json file containing a dashboard definition."+
		"The definition can be retrieved from running sumocli dashboards get or from exporting the dashboard in the UI.")
	cmd.MarkFlagRequired("file")
	return cmd
}

func createDashboard(file string, client *cip.APIClient) {
	var dashboardDefinition types.DashboardRequest
	fileData, err := os.ReadFile(file)
	if err != nil {
		log.Error().Err(err).Msg("failed to read file " + file)
	}
	err = json.Unmarshal(fileData, &dashboardDefinition)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal file")
	}
	data, response, err := client.CreateDashboard(dashboardDefinition)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
