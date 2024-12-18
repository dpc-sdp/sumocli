package pause

import (
	"encoding/json"
	"fmt"
	"github.com/dpc-sdp/sumocli/api"
	"github.com/dpc-sdp/sumocli/pkg/cmd/factory"
	"github.com/dpc-sdp/sumocli/pkg/logging"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdScheduledViewsPause() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "pause",
		Short: "Pause a scheduled view with the given identifier.",
		Run: func(cmd *cobra.Command, args []string) {
			pauseScheduledView(id)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Specify the id of the scheduled view")
	cmd.MarkFlagRequired("id")
	return cmd
}

func pauseScheduledView(id string) {
	var scheduledViewsResponse api.ScheduledViews
	log := logging.GetConsoleLogger()
	requestUrl := "/v1/scheduledViews/" + id + "/pause"
	client, request := factory.NewHttpRequest("POST", requestUrl)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &scheduledViewsResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	scheduledViewsResponseJson, err := json.MarshalIndent(scheduledViewsResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(scheduledViewsResponseJson))
	}
}
